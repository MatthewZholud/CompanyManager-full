package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/server"
)

func main() {
	logger.InitLog()
	server.StartBot()
}


