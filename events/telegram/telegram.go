package telegram

import (
	"main/clients/telegram"
	"main/lib/e/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg: client,
		// offset:  0,
		storage: storage,
	}
}

func (p *Processor) name() {

}
