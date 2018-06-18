package model

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"encoding/xml"
	"github.com/GrigoriiTaranenko/winner_tender/xml_struct/contract"
	base "github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
	"strings"
	_ "github.com/jinzhu/gorm"
	"strconv"
)

// Сведения о контракте
type Contract struct {
	ID               uint      `gorm:"primary_key"`
	SchemeVersion    string
	PublishDate      string
	RegNum           string                         // Номер реестровый записи сведений о контракте. У других PurchaseNumber
	Placing          string                         // Способ размещения заказа 12011 - Электронный аукцион
	IdCustomer       uint                           // ID Заказчика
	IdWinnerSupplier uint                           // ID Победителя поставщика
	IdContractOkpds  string                         // Сортированный по контракту
}

const (
	placingEF = "12011"
	path = "./xml_file/contracts"
)

func GetContractXmlFile() []os.FileInfo{
	log.Println("Пути для files")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return files
}

func (this *Contract) WriteTypeContractXml(name string) (contract.ZfcsContract2015Type, bool) {
	contractXml := contract.ZfcsContract2015Type{}
	regNum :=  strings.Split(name, "_")[1]
	db := base.Db
	db.Where("reg_num = ?", regNum).First(&this)
	if this.ID != 0 {
		log.Println("Документ уже имеется в наличии")
		return contractXml, false
	}
	this.RegNum = regNum
	file, err := os.Open(fmt.Sprintf("%s/%s", path,name))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	contractFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("err = %s", err)
	}
	err = xml.Unmarshal(contractFile, contractXml)
	if err != nil {
		log.Fatalf("err = %s", err)
	}
	return contractXml, true
}

func (this *Contract) Save() {
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

func (this *Contract) FormationModel(contractXml contract.ZfcsContract2015Type, name string) {
	schemeVersion := contractXml.SchemeVersion
	if schemeVersion == "" {
		log.Printf("contractXml = %s", contractXml)
		log.Fatal("Не спарсился документ ", name)
	}
	this.SchemeVersion = schemeVersion
	errorBd := ErrorBd{}
	this.PublishDate  = contractXml.PublishDate.String()
	this.Placing = contractXml.Foundation.FcsOrder.Order.Placing
	if this.Placing != "" {
		message := "Пустой Placing"
		errorBd.Message = message
		log.Println(message)
		return
	}
	if this.Placing != placingEF {
		message := "Placing другого типа " + this.Placing
		errorBd.Message = message
		log.Println(message)
		return
	}

	inn := contractXml.Customer.Inn
	if inn == "" {
		message := "Нету Inn у Заказчика " + name
		errorBd.Message = message
		log.Println(message)
		return
	}
	customer := Customer{}
	if !customer.CheckInn(inn) {
		customer.Name = contractXml.Customer.FullName
		customer.Inn = inn
		customer.Save()
	}
	this.IdCustomer = customer.ID
	supplier := Supplier{}
	b := supplier.AddSupplier(contractXml.Suppliers.Supplier[0])
	if !b {
		message := "Нету Inn у компании" + name
		errorBd.Message = message
		log.Println(message)
		return
	}
	this.IdWinnerSupplier = supplier.ID
	okpdXml := contractXml.Products.Product
	for _, val := range okpdXml {
		 okpd2 := Okpd2{}
		 void := okpd2.AddOkpd(val)
		 if void {
			 message := "Нету Okpd у Заказчика " + name
			 errorBd.Message = message
			 log.Println(message)
			 return
		 }
		 this.IdContractOkpds = fmt.Sprintf("%d ", okpd2.ID)
	}
	this.formationContractOkpds()
	this.Save()
}

func (this *Contract) formationContractOkpds() {
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

func (this *Contract) CountWinnerTender(IdContractOkpds string) (uint) {
	db := base.Db
	rows, err := db.Table("contract").Select("id_winner_supplier, count(id) as countWinner").Where("id_contrackt_okpds = ?", IdContractOkpds).Group("id_contract_okpds").Rows()
	if err != nil {
		log.Fatal(err.Error())
	}
	var count int
//	rowsCountWinner := db.Table("contract").Select("count(id)").Where("id_contrackt_okpds = ?", IdContractOkpds).Row()
	db.Model(&Contract{}).Where("id_contrackt_okpds = ?", IdContractOkpds).Count(&count)
	max := 0
	winnerId := 0
	for rows.Next() {
		columns, _ := rows.Columns()
		max := 0
		winnerId := 0
		for _, column := range columns {
			if count / column["countWinner"] > max {
				max := count / column["countWinner"]
				winnerId = column["id_winner_supplier"]
			}
		}
	}
	return winnerId
}