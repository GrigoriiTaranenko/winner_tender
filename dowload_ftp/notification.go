package dowload_ftp

import (
	"github.com/dutchcoders/goftp"
	"errors"
	"log"
	"net/url"

	"io/ioutil"
	"strings"

	"io"
	"github.com/mihteh/types"
	"github.com/GrigoriiTaranenko/winner_tender/model"
	"strconv"
	"fmt"
)

const NotificationURL = "ftp.zakupki.gov.ru:21"

type  Notification struct {
	FromDate types.Date
	ToDate types.Date
	typeDocument string
	region string
	connect bool
	Size int
	Client *goftp.FTP
	checkDay bool
}

func (this *Notification) NextDirectory(text string) error {
	if !this.connect {
		return errors.New("Нет подключения")
	}
	err := this.Client.Cwd(text)
	if err != nil {
		return err
	}
	cD, _ := this.Client.Pwd()
	log.Printf("текущая директория %s", cD)
	return nil
}

func (this *Notification) GetFromDateString() string {
	return this.FromDate.String()
}

func (this *Notification) GetToDateString() string {
	return this.ToDate.String()
}

func (this *Notification) DownloadFiles() {
	files, err := this.ListClient(this.CurrentDirectory())
	if err != nil  {
		log.Fatalf(err.Error())
	}

	for _, name := range files {

		splitNameZip := strings.Split(name, "_")
		if len(splitNameZip) < 6 {
			continue
		}
		if !CheckBetweenDate(splitNameZip[3], splitNameZip[4], this) {
			continue
		}

		structName := strings.Split(name, " ")
		nameZip := structName[len(structName) - 1]
		nameZip = nameZip[:len(nameZip) - 2]
		downloadedZip := model.DownloadedZip{}
		if downloadedZip.CheckZipFile(nameZip) {
			continue
		}
		this.DownloadFile(nameZip)
	}

}

func (this *Notification) DownloadFile(nameZip string) {
	_, err := this.Client.Retr(nameZip, func(r io.Reader) error {
		read, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatalf("this.Client.Retr %v \n", err.Error())
		}

		err = ioutil.WriteFile(fmt.Sprintf("zip_file/%s", nameZip), read, 0644)

		if err != nil {
			log.Fatalf("ioutil.WriteFile err =  %v \n", err.Error())

		}
		log.Printf("p = %v", read)
		return nil
	})

	if err != nil {
		log.Fatalf("\n %s \n", err.Error())
	}
	downloadedZip := model.DownloadedZip{Name:nameZip}
	downloadedZip.Save()
}

func (this *Notification) CurrentDirectory() string {
	if !this.connect {
		log.Printf("Нет подключения")
		return ""
	}
	dir, _ := this.Client.Pwd()
	return dir
}

func (this *Notification) ListClient(text string) ([]string, error) {
	if !this.connect {
		return nil, errors.New("Нет подключения")
	}
	files, err := this.Client.List(text)
	if err != nil {
		return nil, errors.New(" Ошибка вывода списка клиентов " + err.Error())
	}
	for _, val := range files {
		log.Printf("%v \n", val)
	}
	return files, err
}

func (this *Notification) Connect() error {
	client, err := goftp.Connect(NotificationURL)
	if err != nil {
		return errors.New(" Первая " + err.Error())
	}
	this.Client = client
	err = this.Client.Login("free", "free")
	if err != nil {
		return errors.New(" Вторая " + err.Error())
	}
	log.Println("Подключения успешно установленно")
	this.connect = true
	return nil
}

func (this *Notification) Init(pars url.Values) {
	this.region = pars["region"][0]
	this.typeDocument = pars["typeDocument"][0]

	this.FromDate, _ = types.StringToDate(pars["fromDate"][0])
	this.ToDate, _ = types.StringToDate(pars["toDate"][0])
	err := this.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = this.NextDirectory("fcs_regions")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = this.NextDirectory(this.region)
	if err != nil {
		log.Fatal(err.Error())
	}
	this.NextDirectory(this.typeDocument)
	if err != nil {
		log.Fatal(err.Error())
	}
	dateNow := types.DateNow()
	nowMonth, _ := strconv.Atoi(strings.Split(dateNow.String(), "-")[1])
	toMonth, err := strconv.Atoi(strings.Split(this.ToDate.String(), "-")[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	this.checkDay = false
	if dateNow.Year() != this.ToDate.YearDay() {
		return
	}
	switch nowMonth - toMonth {
	case 0:
		err := this.NextDirectory("currMonth")
		if err != nil {
			log.Fatal(err.Error())
		}
		this.checkDay = true
	case 1:
		err := this.NextDirectory("prevMonth")
		if err != nil {
			log.Fatal(err.Error())
		}
		this.checkDay = true
	}
}