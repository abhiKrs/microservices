package kafka

import (
	logs "web-api/app/utility/logger"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/avro"
)

type RecordsHandler struct {
	Producer *kafka.Producer
	Ser      *avro.GenericSerializer
	Topics   []string
}

func NewRecordHandler(topics []string) (*RecordsHandler, error) {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": BootstrapServersDefault})
	if err != nil {
		logs.WarningLogger.Printf("Failed to create producer: %s\n", err)
		return nil, err
	}

	logs.DebugLogger.Printf("Created Producer %v\n", kafkaProducer)

	schemaClient, err := schemaregistry.NewClient(schemaregistry.NewConfig(SchemaRegistryUrl))
	if err != nil {
		logs.WarningLogger.Printf("Failed to create schema registry client: %s\n", err)
		return nil, err
	}

	ser, err := avro.NewGenericSerializer(schemaClient, serde.ValueSerde, avro.NewSerializerConfig())

	if err != nil {
		logs.WarningLogger.Printf("Failed to create serializer: %s\n", err)
		return nil, err
	}

	rh := &RecordsHandler{
		Producer: kafkaProducer,
		Ser:      ser,
		Topics:   topics,
	}
	return rh, nil
}
