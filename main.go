//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"hello-wire/handler"
	"hello-wire/model"
	"hello-wire/service"
	"log"
)

func initMailHandler() (*handler.MailHandler, error) {
	wire.Build(handler.MailHandlerSet)

	return &handler.MailHandler{}, nil
}

func initMailHandlerMocked(_ service.MailStore) (*handler.MailHandler, error) {
	wire.Build(handler.MailHandlerMockedSet)

	return &handler.MailHandler{}, nil
}

func main() {
	mailHandler, err := initMailHandler()
	if err != nil {
		log.Fatal(err)
	}

	mails, err := mailHandler.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("mails: %+v\n", mails)

	mailHandlerMocked, err := initMailHandlerMocked(NewMailStoreMock())
	mails, err = mailHandlerMocked.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("mocked mails: %+v\n", mails)
}

type MailStoreMock struct{}

func (m *MailStoreMock) GetAll() ([]model.Mail, error) {
	return []model.Mail{}, nil
}

func NewMailStoreMock() service.MailStore {
	return &MailStoreMock{}
}
