package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot/server"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/interService"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/redis"
)

func main() {
	logger.InitLog()
	redis := redis.Initialize()
	kafka := kafka.Initialize()
	newBot := server.StartBot()

	interService := interService.Initialize(kafka)

	updates := bot.NewUpdateChan(newBot.BotAPI, redis, interService)
	updates.Listen()
}
