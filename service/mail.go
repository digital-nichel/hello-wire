package service

import (
	"github.com/google/wire"
	"hello-wire/logger"
	"hello-wire/model"
	"hello-wire/store"
)

type AccountStore interface {
	GetAll() ([]model.Account, error)
}

type MailStore interface {
	GetAll() ([]model.Mail, error)
}

type MailService struct {
	logger       logger.Logger
	accountStore AccountStore
	mailStore    MailStore
}

func ProvideMailService(logger logger.Logger, accountStore AccountStore, mailStore MailStore) *MailService {
	return &MailService{
		logger:       logger,
		accountStore: accountStore,
		mailStore:    mailStore,
	}
}

func (service *MailService) GetAll() ([]model.Mail, error) {
	return service.mailStore.GetAll()
}

var MailServiceMockedSet = wire.NewSet(
	ProvideMailService,

	logger.LoggerSet,

	store.AccountStoreSet,
	wire.Bind(new(AccountStore), new(*store.AccountStore)),

	//store.MailStoreSet,
	//wire.Bind(new(MailStore), new(*store.MailStore)),
)

var MailServiceSet = wire.NewSet(
	MailServiceMockedSet,

	store.MailStoreSet,
	wire.Bind(new(MailStore), new(*store.MailStore)),
)
