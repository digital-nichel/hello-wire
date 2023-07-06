package store

import (
	"github.com/google/wire"
	"hello-wire/model"
)

type MailStore struct {
	mails []model.Mail
}

func ProvideMailStore() *MailStore {
	return &MailStore{
		mails: []model.Mail{
			{
				Id:    "1",
				Title: "title1",
				Body:  "body1",
			},
		},
	}
}

func (store *MailStore) GetAll() ([]model.Mail, error) {
	return store.mails, nil
}

var MailStoreSet = wire.NewSet(
	ProvideMailStore,
)
