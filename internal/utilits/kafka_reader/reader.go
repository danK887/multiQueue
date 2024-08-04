package kafkareader

import (
	"encoding/json"
	"fmt"
	"log"
	"multiQueue/internal/config"

	"github.com/IBM/sarama"
)

func ReadMsgKafka(partConsumer sarama.PartitionConsumer) (config.Message, error) {
	log.Println("Read message from kafka")
	message := <-partConsumer.Messages()

	// Декодирование сообщения в структуру config.Message
	var msg config.Message
	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		return config.Message{}, fmt.Errorf("failed to unmarshal message: %w", err)
	}

	return msg, nil
}
