package account

import (
	"testing"
	"time"
)

// test to set the title
func TestAccount_SetTitle(t *testing.T) {

	newAccount := Account{}

	newAccount.SetTitle("TestTitle");

	if newAccount.Title  != "TestTitle" {
		t.Fail();
	}

}

// test to set the title
func TestAccount_SetBeginDate(t *testing.T) {

	newAccount := Account{}

	//set begin date
	newAccount.SetBeginDate("05.04.2013");

	i, _ := time.Parse("02.01.2006", "05.04.2013")
	if newAccount.BeginDate  != i {
		t.Fail();
	}

}

//test to set the end date
func TestAccount_SetEndDate(t *testing.T) {

	newAccount := Account{}

	//set begin date
	newAccount.SetEndDate("05.04.2016")

	i, _ := time.Parse("02.01.2006", "05.04.2016")
	if newAccount.EndDate  != i {
		t.Fail();
	}
}

//test the SetAccountBalance method and convert it to float (if string)
func TestAccount_SetAccountBalance(t *testing.T) {

	newAccount := Account{}

	result, err := newAccount.SetBalance("3.333,77")

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Error(err)
	}

	if newAccount.Balance != float64(3333.77)  {
		t.Fail()
	}

}

//test the SetAccountBalance method and convert it to float (if string)
func TestAccount_SetAccountBalanceBig(t *testing.T) {

	newAccount := Account{}

	result, err := newAccount.SetBalance("413.333,77")

	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Error(err)
	}

	if newAccount.Balance != float64(413333.77)  {
		t.Fail()
	}

}

