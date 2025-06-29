package main

import (
	"flag"
	"log"
	"main/clients/telegram"
)

const (
	hostTg = "api.telegram.org"
)

const (
	host = "localhost"
)

func main() {
	tgClient := telegram.New(hostTg, mustToken())
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
