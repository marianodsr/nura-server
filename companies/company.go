package companies

import (
	"github.com/marianodsr/nura-server/storage"
	"github.com/marianodsr/nura-server/users"
)

type Company struct {
	storage.Base
	Name         string
	Logo         string
	AllowedUsers []users.User
}
