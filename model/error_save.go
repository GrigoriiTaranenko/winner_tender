package model

import (
	"log"
	"github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
)

type ErrorBd struct {
	ID              uint      `gorm:"primary_key"`
	Message         string
}

func (this *ErrorBd) Save() {
	db := connect_base.Db
	if this.ID == 0 {
		db.Create(&this)
	} else {
		db.Save(&this)
	}
	if db.Error != nil {
		log.Fatalf("db.Error = %s", db.Error)
	}
}