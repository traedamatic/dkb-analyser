package account

import (
	"testing"
	"time"
)

// test to set the title
func TestSetTitle(t *testing.T) {

	newAccount := Account{}

	newAccount.SetTitle("TestTitle");

	if newAccount.Title  != "TestTitle" {
		t.Fail();
	}

}

// test to set the title
func TestSeBeginDate(t *testing.T) {

	newAccount := Account{}

	//set begin date
	newAccount.SetBeginDate("05.04.2013");

	i, _ := time.Parse("02.01.2006", "05.04.2013")
	if newAccount.BeginDate  != i {
		t.Fail();
	}

}

