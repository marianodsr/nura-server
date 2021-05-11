package storage

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Base struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4()
	base.ID = uuid.String()
	return nil

}

func InitDB() error {
	dsn := "host=localhost user=postgres password=auth159123456 dbname=exo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func GetDBConnection() *gorm.DB {
	return db
}
