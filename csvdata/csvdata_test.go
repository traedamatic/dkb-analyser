package csvdata

import (
	"testing"
)

//define stub account struct for data
type TestAccount struct {
}

func (a *TestAccount) SetTitle(title string) (bool, error) {
	return true, nil
}

func (a *TestAccount) SetBeginDate(stringDate string) (bool, error) {
	return true, nil
}

func (a *TestAccount) SetEndDate(stringDate string) (bool, error) {
	return true, nil
}

func (a *TestAccount) SetBalance(stringDate string) (bool, error) {
	return true, nil
}


//read data and set the account title
func TestParseDataAndSetAccountTitle(t *testing.T) {

	csvData := []string{"Kontonummer:", "DE666668888888 / Private"}

	newAccount := &TestAccount{}

	result, err := ParseCSV(newAccount, csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}

//test empty slice of csv data
func TestReadDataWithEmptySlice(t *testing.T) {

	csvData := []string{}

	newAccount := &TestAccount{}

	result, err := ParseCSV(newAccount, csvData)

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

	newAccount := &TestAccount{}
	result, err := ParseCSV(newAccount, csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}

func TestParseCSV_EndDate(t *testing.T) {

	csvData := []string{"Bis:", "05.04.2016"}

	newAccount := &TestAccount{}
	result, err := ParseCSV(newAccount, csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}

func TestParseCSV_CurrentBalance(t *testing.T) {

	csvData := []string{"Kontostand vom:", "9.999,45"}

	newAccount := &TestAccount{}
	result, err := ParseCSV(newAccount, csvData)

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}