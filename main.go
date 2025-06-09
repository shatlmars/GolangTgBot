package main

import (
	"flag"
	"log"
	"main/clients/telegram"
)

func main() {

	// fmt.Println(t)
	tgClient := telegram.New("api.telegram.org", mustToken())
	// token = flags.Get(token)
	//fetcher = fetcher.New()
	//consumer.Start()
	//
}

func mustToken() string {

	//bot -tg-bot-token 'my token'
	token := flag.String("token-bot-token",
		"",
		"token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
