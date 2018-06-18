package model

import (
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"log"
)

type DownloadedZip struct {
	ID   uint    `gorm:"primary_key"`
	Name string  `gorm:"unique_index;not null"`
}

func (this *DownloadedZip) Save() {
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

func (this *DownloadedZip) CheckZipFile(name string) bool {
	base.Db.Where("name = ?", name).First(&this)
	if this.ID != 0 {
		return true
	}
	return false
}