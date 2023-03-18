package kafka

import (
	"encoding/json"

	logs "web-api/app/utility/logger"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func SendLog(kafkaTopic string, payload interface{}) error {
	// Produce a new record to the topic...
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServersDefault,
	})

	if err != nil {
		logs.ErrorLogger.Printf("Failed to create producer: %s", err)
		return err
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		logs.ErrorLogger.Println(err)
		return err
	}

	producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &kafkaTopic,
				Partition: 0,
			},
			Value: bytePayload,
		},
		nil,
	)

	// Wait for delivery report
	e := <-producer.Events()

	message := e.(*kafka.Message)
	if message.TopicPartition.Error != nil {
		logs.WarningLogger.Printf("failed to deliver message: %v\n", message.TopicPartition)
	} else {
		logs.WarningLogger.Printf("delivered to topic %s [%d] at offset %v\n",
			*message.TopicPartition.Topic,
			message.TopicPartition.Partition,
			message.TopicPartition.Offset)
	}

	producer.Close()
	return nil
}
