package dowload_ftp

import (
	"time"
	"github.com/jlaffaye/ftp"
	"errors"
	"log"
)

const  NotificationURL = "ftp.zakupki.gov.ru:21"

type  Notification struct {
	FromDate time.Time
	ToDate time.Time
	connect bool
	Size int
	Client *ftp.ServerConn
}

func (this *Notification) NextDirectory(text string) error {
	if !this.connect {
		return errors.New("Нет подключения")
	}
	err := this.Client.ChangeDir(text)
	if err != nil {
		return err
	}
	cD, _ := this.Client.CurrentDir()
	log.Printf("текущая директория %s", cD)
	return nil
}

func (this *Notification) DownloadFile() {
	
}

func (this *Notification) CurrentDirectory() string {
	if !this.connect {
		log.Printf("Нет подключения")
		return ""
	}
	dir, _ := this.Client.CurrentDir()
	return dir
}

func (this *Notification) ListClient(text string) error {
	if !this.connect {
		return errors.New("Нет подключения")
	}
	entries, err := this.Client.List(text)
	if err != nil {
		return errors.New(" Ошибка вывода списка клиентов " + err.Error())
	}
	for _, val := range entries {
		log.Printf("%v \n", val)

	}
	return nil
}

func (this *Notification) Connect() error {
	client, err := ftp.Dial(NotificationURL)
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