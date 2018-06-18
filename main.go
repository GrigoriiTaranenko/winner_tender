package main

import (
	"os"
	"fmt"
	"log"
	"html/template"

	"net/http"

	"io/ioutil"
	"github.com/julienschmidt/httprouter"

	"github.com/GrigoriiTaranenko/winner_tender/model"
	"github.com/GrigoriiTaranenko/winner_tender/pars_zip"
	"github.com/GrigoriiTaranenko/winner_tender/dowload_ftp"
	"github.com/GrigoriiTaranenko/connect_postgresql/connect_base"
)

type errorType struct {
	errorText string
}

func outputError(err error) {
	if err == nil {
		return
	}
	log.Printf("err = %s \n", err.Error())
}

func htmlError(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}
	errorHtml := errorType{
		errorText: err.Error(),
	}
	tmpl, _ := template.ParseFiles("client/error.html")
	tmpl.Execute(w, errorHtml)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, _ := template.ParseFiles("client/index.html")
	tmpl.Execute(w, nil)
}

func xmlDownload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pars := r.URL.Query()

	err := dowload_ftp.XmlForm(pars)
	if err != nil {
		htmlError(err, w)
		return
	}

	download := &dowload_ftp.Notification{}

	download.Init(pars)
	log.Println("Данные иннициализированны")
	download.DownloadFiles()
	log.Println("Данные загруженны")
	http.Redirect(w, r, "../", http.StatusSeeOther)
}

func ParsingZip(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files, err := ioutil.ReadDir("./zip_file")
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range files {
		log.Printf("val = %s", val.Name())
		pars_zip.DirectionUnpacking(val.Name())
		err := os.Remove(fmt.Sprintf("zip_file/%s", val.Name()))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	log.Println("Парсинг успешно завершился")
	//http.Redirect(w, r, "../", http.StatusSeeOther)
}

func ParsingContract(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files := model.GetContractXmlFile()
	for _, f := range files {
		nameF := f.Name()
		log.Println(nameF)
		contractModel := model.Contract{}
		contractXml, bPars := contractModel.WriteTypeContractXml(nameF)
		if !bPars {
			continue
		}
		contractModel.FormationModel(contractXml, nameF)

	}
}

func WinnerTender(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pars := r.URL.Query()
	regNum := pars["regNum"][0]
	notification := model.FcsNotificationEf{}
	if !notification.CheckRegNum(regNum) {
		log.Printf("Данного уведомления нету")
	}
	contract := model.Contract{}
	winnerSupplierId := contract.CountWinnerTender(notification.IdContractOkpds)
	supplier := model.Supplier{}
	supplier.Find(winnerSupplierId)
	log.Printf("Победитель тендера %s скорее всего будет %s", regNum, supplier.Name)
	
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	connect_base.ConnectToDB()
	db := connect_base.Db
	defer db.Close()

	db.AutoMigrate(&model.DownloadedZip{})
	db.AutoMigrate(&model.SparseFile{})
	db.AutoMigrate(&model.Supplier{})
	db.AutoMigrate(&model.ProtocolEF3{})
	db.AutoMigrate(&model.SupplierEF3{})
	db.AutoMigrate(&model.Okpd2{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Contract{})
	db.AutoMigrate(&model.ErrorBd{})
	db.AutoMigrate(&model.FcsNotificationEf{})

	connect_base.ConnectToDB()
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/download", xmlDownload)
	router.GET("/parsing", ParsingZip)
	router.GET("/contract", ParsingContract)
	router.GET("/winner", WinnerTender)

	log.Fatal(http.ListenAndServe(":8080", router))



/*	notification := &dowload_ftp.Notification{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		switch command {
		case "connect":
			outputError(notification.Connect())
		case "curDir":
			log.Printf("%s \n", notification.CurrentDirectory())
		case "listData":
			outputError(notification.ListClient(notification.CurrentDirectory()))
		case "download":
			outputError(notification.DownloadFile())
		default:
			outputError(notification.NextDirectory(command))
		}
	}
	log.Println("The end")*/
}