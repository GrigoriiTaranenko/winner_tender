package main

import (
	"github.com/GrigoriiTaranenko/winner_tender/dowload_ftp"
	"log"
)

func main()  {
	notification := &dowload_ftp.Notification{}
	err := notification.Connect()
	if err != nil {
		log.Printf("err = %s \n", err.Error())
	}
	log.Println("The end")
}