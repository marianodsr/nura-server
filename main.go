package main

import (
	"fmt"

	"github.com/marianodsr/nura-server/storage"
	"github.com/marianodsr/nura-server/users"
)

func main() {
	fmt.Println("Exo app!")
	if err := storage.InitDB(); err != nil {
		panic(err)
	}
	db := storage.GetDBConnection()
	db.AutoMigrate(&users.User{})
	db.Create(&users.User{
		FirstName: "test",
		LastName:  "test2",
		Email:     "test@gmail.com",
		Password:  "12312321",
	})
	registerRoutes()
}
