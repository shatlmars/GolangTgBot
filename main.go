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
	tgClient := telegram.New(hostTg, mustToken())
	//fetcher = fetcher.NEW
	//proccessor = proccessor.New()
	//consumer = consumer.New(fetcher, proccessor)
	tgClient.Updates(1, 1)

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
