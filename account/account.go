package account

import (
	"time"
)

type Account struct {
	Title string
	BeginDate time.Time
}

// read data from a string array and add it to the account struct
func (a *Account) ReadData(data []string) (bool, error) {

	if len(data) == 0 {
		return false, nil
	}

	//check the first entry in the string array
	switch data[0] {

	case "Kontonummer:" :
		a.Title = data[1]
		return true, nil

	case "Von:":
		date, parseError := time.Parse("02.01.2006", data[1])

		if parseError != nil {
			return false, parseError
		}

		a.BeginDate = date

		return true, nil
	}

	return false, nil
}