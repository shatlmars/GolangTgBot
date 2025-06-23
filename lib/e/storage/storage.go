package storage

import "crypto/sha1"

type Storage interface {
	Save(p *Page) error
	PickRandom(username string) (*Page, error)
	Remove(p *Page) error
	IsExsist(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	// Created time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
}
