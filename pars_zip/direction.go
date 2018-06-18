package pars_zip

import (
	"archive/zip"
	"fmt"
	"log"
	"strings"
	"io"
	"os"
)


var parsZip = map[string][]string{
	"protocol": {
		"fcsProtocolEF3",
	},
	"contract": {
		"contract",
	},
	"notification":{
		"fcsNotificationEF",
		"fcsNotificationEA44",
	},
}

func IndexOf(stringArray []string, str string) int {
	for i, val := range stringArray {
		if val == str {
			return i
		}
	}
	return -1
}

// Возвращает путь хранения
func DirectionUnpacking(name string) {
	r, err := zip.OpenReader(fmt.Sprintf("zip_file/%s", name))

	if err != nil {
		log.Fatal(err.Error())
	}
	defer r.Close()
//	pathSaveFile := strings.Split(name, "_")[0]
	for _, f := range r.File {
		formatFile := strings.Split(f.Name, ".")[1]
		typeFile := strings.Split(f.Name, "_")[0]
		typeZip := strings.Split(name, "_")[0]
		if formatFile == "sig" || IndexOf(parsZip[typeZip], typeFile) == -1 {
			continue
		}
		log.Printf("fileNmae = %s", f.Name)
		rc, err := f.Open()
		file, err := os.Create(fmt.Sprintf("xml_file/%ss/%s", typeZip, f.Name))
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(file, rc)
		if err != nil {
			log.Fatal(err)
		}

		file.Close()
		rc.Close()
	}

}
