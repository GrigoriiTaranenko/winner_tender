package model
// Схема электронного аукциона третиго второго этапа
import (
	"time"
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"log"
)

type ProtocolEF3 struct {
	ID             uint      `gorm:"primary_key"`
	SchemeVersion  string
	PublishDate    time.Time                                 // Дата публикации заявки
	PurchaseNumber string    `gorm:"unique_index;not null"`  // Номер закупки
	IdCustomer     int                                       // Id Заказчика может быть пустым
	SupplierEF3    []SupplierEF3
}

func (this *ProtocolEF3) Save() {
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

