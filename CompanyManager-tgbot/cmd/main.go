package main

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
)

func main() {
	logger.InitLog()
	internal.StartBot()
}


