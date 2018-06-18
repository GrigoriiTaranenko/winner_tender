package model

// Хранит данные по определенному
import (
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"log"
)

type SupplierEF3 struct {
	ID               uint   `gorm:"primary_key"`
	IdSupplier       uint                        // Внешний ключ поставщика
	IdFcsProtocolEF3 uint   	                 // Внешний ключ протокола
	WinnerPrice      uint                        // Цена поставки
	AppRating        uint                        // Рейтинг по списку
}

func (this *SupplierEF3) Save() {
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