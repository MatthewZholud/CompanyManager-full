package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/handlers"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/redis"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/server"
)

func main() {
	logger.InitLog()
	redis := redis.Initialize()
	bot := server.StartBot()
	updates := handlers.NewUpdateChan(bot.BotAPI, redis)
		updates.Listen()
}


