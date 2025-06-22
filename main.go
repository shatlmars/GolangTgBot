package main

import (
	"flag"
	"log"
	"main/clients/telegram"
)

const (
	hostTg = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(hostTg, mustToken())

	//fetcher = fethcer.New()
	//proccessor = proccessof.New()

	//consumer.Start(fetcher, proccessor)
	// tgClient.SendMessage()
}

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
