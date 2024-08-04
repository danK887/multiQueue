package kafkamsgwrite

import (
	"encoding/json"
	"log"
	"multiQueue/internal/config"

	"github.com/IBM/sarama"
	"github.com/Pallinder/go-randomdata"
)

func MsgWrite(producer sarama.SyncProducer, topic string) {

	go func() {
		msg := config.Message{
			Name:   randomdata.LastName(),
			Age:    randomdata.Number(18, 65),
			Phone:  randomdata.PhoneNumber(),
			Salary: randomdata.Decimal(55000, 350000, 2),
		}
		jsonMessage, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("json marshal error: %v", err)
		}
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(jsonMessage),
		}
		_, _, err = producer.SendMessage(message)
		if err != nil {
			log.Fatalf("send message error: %v", err)
		}
	}()
}
