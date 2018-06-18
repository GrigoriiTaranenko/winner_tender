package model

import (
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"log"
	"strings"
	"strconv"
)

type FcsNotificationEf struct {
	ID              uint      `gorm:"primary_key"`
	SchemeVersion   string
	PublishDate     string
	Placing         string                         // Способ размещения заказа 12011 - Электронный аукцион
	RegNum          string                         // Номер реестровый записи сведений о контракте. У других PurchaseNumber
	IdCustomer      uint                           // ID Заказчика
	IdContractOkpds  string                        // Сортированный по контракту
}

func (this *FcsNotificationEf) Save() {
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

func (this *FcsNotificationEf) CheckRegNum(regNum string) bool {
	base.Db.Where("regNum = ?", regNum).First(&this)
	if this.ID != 0 {
		return true
	}
	return false
}

func (this *FcsNotificationEf) formationContractOkpds() {
	this.IdContractOkpds = strings.TrimSpace(this.IdContractOkpds)
	okpds := strings.Split(this.IdContractOkpds, " ")
	count := len(okpds)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			el, err := strconv.Atoi(okpds[i])
			if err != nil {
				log.Fatal(err.Error())
			}
			el1, err := strconv.Atoi(okpds[j])
			if err != nil {
				log.Fatal(err.Error())
			}
			if el > el1 {
				k := okpds[i]
				okpds[i] = okpds[j]
				okpds[j] = k
			}
		}
	}
	this.IdContractOkpds = strings.Join(okpds, " ")
}