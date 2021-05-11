package users

import (
	"github.com/marianodsr/nura-server/companies"
	"github.com/marianodsr/nura-server/storage"
)

type User struct {
	storage.Base
	Email     string
	Password  string
	FirstName string
	LastName  string
	Companies []companies.Company
}
