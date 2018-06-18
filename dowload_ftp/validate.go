package dowload_ftp

import (
	"net/url"
	"github.com/pkg/errors"
	"github.com/mihteh/types"
	"strings"
)

type BetweenDate interface {
	GetFromDateString() string
	GetToDateString() string
}

func XmlForm(pars url.Values) error {
	if pars["region"][0] == "" {
		return errors.New("Не заполненна форма региона")
	}
	if pars["typeDocument"][0] == "" {
		return errors.New("Не заполнена форма типа документа")
	}
	fromDate := pars["fromDate"][0]
	if fromDate == "" {
		return errors.New("Не заполнена дата начала выборки")
	}
	toDate := pars["toDate"][0]
	if toDate == "" {
		return errors.New("Не заполнена дата окончания выборки")
	}
	dateNow := types.DateNow().String()
	if dateNow < toDate || dateNow < fromDate || fromDate > toDate {
		return errors.New("Ошибка работы с датами")
	}
	return nil
}

func CheckBetweenDate(fromDateFile string, toDateFile string, datesLoad BetweenDate) bool {
	fromDate := strings.Replace(datesLoad.GetFromDateString(), "-", "", -1)
	toDate := strings.Replace(datesLoad.GetToDateString(), "-", "", -1)
	if (fromDate >= fromDateFile && fromDate <= toDateFile) || (toDate >= fromDateFile && toDate <= toDateFile){
		return true
	}
	return false
}