package model

import (
	"log"
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
)


// Справочник заказчиков.
type Customer struct {
	ID                uint              `gorm:"primary_key"`
	Name              string            `gorm:"not null"`
	Inn               string            `gorm:"unique_index;not null"`
	Contract          Contract
	FcsNotificationEf FcsNotificationEf
}

func (this *Customer) Save() {
	db := base.Db
	if this.ID == 0 {
		db.Create(&this)
	} else {
		db.Save(&this)
	}
	if db.Error != nil {
		log.Fatalf("db.Error = %s", db.Error)
	}
}

func (this *Customer) CheckInn(inn string) bool {
	base.Db.Where("inn = ?", inn).First(&this)
	if this.ID != 0 {
		return true
	}
	return false
}