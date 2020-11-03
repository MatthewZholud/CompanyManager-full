package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot/botServer"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/bot/updateListener"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/interService"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/kafka"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/redis"
)

func main() {
	logger.InitLog()
	redis := redis.StartRedis()
	newBot := botServer.StartBot()
	kafka := kafka.Initialize()


	interService := interService.Initialize(kafka)

	botHandlerService := handlers.NewHandlerService(newBot.BotAPI, redis, interService)

	updates := updateListener.NewUpdateChan(newBot.BotAPI)
	updates.Listen(botHandlerService)
}
