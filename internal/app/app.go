package app

import (
	"log/slog"
	connectkafka "multiQueue/internal/connect_kafka"
	writer "multiQueue/internal/utilits/kafka_msg_write"
	reader "multiQueue/internal/utilits/kafka_reader"
)

func Run() {
	topic := "multiQueue"
	producer, consumer, partitionConsumer := connectkafka.Connect_KK("10.130.2.30:9093", topic)

	slog.Info("Start application")
	writer.MsgWrite(producer, topic)
	slog.Info("Message was send to kafka")

	// прочитать сообщения из кафки
	reader.ReadMsgKafka(partitionConsumer)

	// подключиться к ребиту

	// записать прочитанные сообщения из кафки в ребит

	// подлкючиться к кликхаус

	// из ребита прочитать сообщения и записать в кликхаус

}
