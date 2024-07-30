package connectkafka

import (
	"log"

	"github.com/IBM/sarama"
)

// Connect создает подключение к Kafka и возвращает синхронного продюсера.
func Connect_KK(broker, topic string) (sarama.SyncProducer, sarama.Consumer, sarama.PartitionConsumer) {

	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create sync producer: %v", err)
	}

	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}

	return producer, consumer, partitionConsumer
}
