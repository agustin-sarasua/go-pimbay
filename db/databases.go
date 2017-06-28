package db

import "github.com/agustin-sarasua/pimbay/model"

type UserDatabase interface {
	SaveUser(u *model.User) (id int64, e error)
	GetUser(id int64) (*model.User, error)
	Close()
}

type AccountDatabase interface {
	SaveAccount(a *model.Account) (id int64, e error)
	GetAccount(id int64) (*model.Account, error)
}
