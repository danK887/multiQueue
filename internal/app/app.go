package app

import (
	"encoding/json"
	"log"
	connectkafka "multiQueue/internal/connect_kafka"

	"github.com/IBM/sarama"
	"github.com/Pallinder/go-randomdata"
)

func Run() {

	topic := "multiQueue"
	producer, _, _ := connectkafka.Connect_KK("10.130.2.30:9093", topic)
	defer producer.Close()

	type Message struct {
		Name   string  `json:"name"`
		age    int     `json:"age"`
		Phone  string  `json:"phone"`
		Salary float64 `json:"salary"`
	}

	go func() {
		msg := Message{
			Name:   randomdata.LastName(),
			age:    randomdata.Number(18, 65),
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
