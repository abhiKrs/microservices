package kafka

import (
	"context"
	"time"
	logs "web-api/app/utility/logger"

	// "github.com/Shopify/sarama"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateTopic(topic string) error {

	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServersDefault,
	})

	if err != nil {
		logs.ErrorLogger.Printf("Failed to create Admin client: %s\n", err)
		return err
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDuration, err := time.ParseDuration("60s")
	if err != nil {
		panic("time.ParseDuration(60s)")
	}

	// Check kafka deployment options----------------------------------
	results, err := adminClient.CreateTopics(ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 3}},
		kafka.SetAdminOperationTimeout(maxDuration))

	if err != nil {
		logs.ErrorLogger.Printf("Problem during the topic creation: %v\n", err)
		return err
	}

	// Check for specific topic errors
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError &&
			result.Error.Code() != kafka.ErrTopicAlreadyExists {
			logs.ErrorLogger.Printf("Topic creation failed for %s: %v",
				result.Topic, result.Error.String())
			return err
		}
	}

	adminClient.Close()
	return nil

}
