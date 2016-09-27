package account

import "testing"

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
