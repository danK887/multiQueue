package app

import (
	"log/slog"
	writer "multiQueue/internal/utilits/kafka_msg_write"
)

func Run() {
	slog.Info("Start application")
	writer.MsgWrite()
	slog.Info("Message was send to kafka")

	// прочитать сообщения из кафки

	// подключиться к ребиту

	// записать прочитанные сообщения из кафки в ребит

	// подлкючиться к кликхаус

	// из ребита прочитать сообщения и записать в кликхаус

}
