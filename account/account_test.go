package account

import (
	"testing"
	"time"
)

//read data and set the account title
func TestParseDataAndSetAccountTitle(t *testing.T) {

	csvData := []string{"Kontonummer:", "DE666668888888 / Private"}

	newAccount := Account{}

	result, err := newAccount.ReadData(csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	if newAccount.Title != "DE666668888888 / Private" {
		t.Fail()
	}

}

//test empty slice of csv data
func TestReadDataWithEmptySlice(t *testing.T) {

	csvData := []string{}

	newAccount := Account{}

	result, err := newAccount.ReadData(csvData)

	if result != false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}

//test beginning date
func TestReadDataFromDate(t *testing.T) {

	csvData := []string{"Von:", "05.04.2013"}

	newAccount := Account{}

	result, err := newAccount.ReadData(csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	i, _ := time.Parse("02.01.2006", "05.04.2013")
	if newAccount.BeginDate != i {
		t.Fail()
	}

}
