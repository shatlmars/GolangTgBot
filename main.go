package main

import (
	"errors"
	"flag"
	"log"
)

func main() {

	// // fmt.Println(t)
	// tgClient := telegram.New("api.telegram.org", "8096732712:AAFwSHFTtjK9_-f-Zlr1GR_slPGsvCGmMdU")
	// // token = flags.Get(token)
	// //fetcher = fetcher.New()
	// //consumer.Start()
	// //

	// tgClient.SendMessage(295673061, "Hello")
	err := errors.New("sd")
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
