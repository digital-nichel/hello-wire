package store

import (
	"github.com/google/wire"
	"hello-wire/model"
)

type AccountStore struct {
	accounts []model.Account
}

func ProvideAccountStore() *AccountStore {
	return &AccountStore{
		accounts: []model.Account{
			{
				Id:       "1",
				Email:    "email1",
				Password: "password1",
			},
		},
	}
}

func (store *AccountStore) GetAll() ([]model.Account, error) {
	return store.accounts, nil
}

var AccountStoreSet = wire.NewSet(
	ProvideAccountStore,
)
