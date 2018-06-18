package model

import (
	"log"
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"github.com/GrigoriiTaranenko/winner_tender/xml_struct/contract"

)

// Справочник поставщиков
type Supplier struct {
	ID          uint           `gorm:"primary_key"`
	Name        string         `gorm:"not null"`
	Inn         string         `gorm:"unique_index;not null"`
	SupplierEF3 []SupplierEF3
	Contract    []Contract
}

func (this *Supplier) Save() {
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

func (this *Supplier) CheckInn(inn string) bool {
	base.Db.Where("inn = ?", inn).First(&this)
	if this.ID != 0 {
		return true
	}
	return false
}

func (this *Supplier) AddSupplier(supplier contract.ZfcsContract2015SupplierType) bool {
	if supplier.LegalEntityRF.INN != "" {
		this.add(supplier.LegalEntityRF.INN, supplier.LegalEntityRF.FullName)
		return true
	}
	if supplier.IndividualPersonRF.INN != "" {
		this.add(supplier.IndividualPersonRF.INN, supplier.IndividualPersonRF.FirstName)
		return true
	}
	if supplier.IndividualPersonRFisCulture.INN != "" {
		this.add(supplier.IndividualPersonRFisCulture.INN, supplier.IndividualPersonRFisCulture.FirstName)
		return true
	}
	return false
}

func (this *Supplier) add(inn string, name string) {
	if this.CheckInn(inn) {
		return
	}
	this.Name = name
	this.Inn = inn
}

func (this *Supplier) Find(id uint) {
	base.Db.Where("id = ?", id).First(&this)
}
