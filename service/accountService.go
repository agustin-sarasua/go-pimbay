package service

import (
	"log"

	"github.com/agustin-sarasua/pimbay/model"
	mgo "gopkg.in/mgo.v2"
)

const (
	accountCollection = "account"
)

func SaveAccount(s *mgo.Session, a *model.Account) <-chan bool {
	out := make(chan bool)
	//defer finally()
	go func() {
		session := s.Copy()
		defer session.Close()

		c := session.DB(dbName).C(accountCollection)
		err := c.Insert(a)
		if err != nil {
			if mgo.IsDup(err) {
				panic(err)
			}
			log.Println("Failed insert User: ", err)
			out <- false
		} else {
			out <- true
		}
	}()
	return out
}
