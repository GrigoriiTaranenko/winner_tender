package dowload_ftp

import (
	"time"
	"github.com/jlaffaye/ftp"
	"fmt"
	"errors"
)

const  NotificationURL = "ftp.zakupki.gov.ru:21"

type  Notification struct {
	FromDate time.Time
	ToDate time.Time
	Size int
}

func (this *Notification) Connect() error {
	client, err := ftp.Dial(NotificationURL)
	if err != nil {
		return errors.New(" Первая " + err.Error())
	}
	err = client.Login("free", "free")
	if err != nil {
		return errors.New(" Вторая " + err.Error())
	}
	entries, err := client.List("/")
	if err != nil {
		return errors.New(" Третья " + err.Error())
	}
	fmt.Printf("entries = %s", entries)
	return nil

	/*	client := &http.Client{
			Timeout: 60 * time.Second,
		}
		transport := &http.Transport{}
		transport.RegisterProtocol("")
		resp, err := client.Get(NotificationURL)
		if err != nil {
			log.Printf("resp err = %s \n", err.Error())
			return
		}
		log.Printf("body = %s \n", resp.Body)*/
}