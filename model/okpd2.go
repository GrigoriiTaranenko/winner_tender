package model

import (
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"github.com/GrigoriiTaranenko/winner_tender/xml_struct/contract"
	"log"
)

type Okpd2 struct {
	ID             uint   `gorm:"primary_key"`
	Number         uint
	Name           string                       // Имя закупки
	Code           string                       // Код закупки
}

func (this *Okpd2) Save() {
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

func (this *Okpd2) CheckCode(code string, number uint) bool {
	base.Db.Where("code = ? and number = ?", code, number).First(&this)
	if this.ID != 0 {
		return true
	}
	return false
}

func (this *Okpd2) AddOkpd(product contract.ProductsProduct) bool {
	if product.OKPD.Code != "" {
		Okpd := product.OKPD
		this.add(Okpd.Code, Okpd.Name, 1)
		return true
	}
	if product.OKPD2.Code != "" {
		Okpd2 := product.OKPD2
		this.add(Okpd2.Code, Okpd2.Name, 2)
		return true
	}
	return false
}

func (this *Okpd2) add(code string, name string, number uint) {
	if this.CheckCode(code, number) {
		return
	}
	this.Code = code
	this.Name = name
	this.Number = number
	this.Save()
}