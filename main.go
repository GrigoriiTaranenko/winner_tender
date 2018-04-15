package main

import (
	"github.com/GrigoriiTaranenko/winner_tender/dowload_ftp"
	"log"
	"bufio"

	"os"
)

func outputError (err error) {
	if err == nil {
		return
	}
	log.Printf("err = %s \n", err.Error())
}
func main()  {
	notification := &dowload_ftp.Notification{}
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
		default:
			outputError(notification.NextDirectory(command))
		}
	}
	log.Println("The end")
}