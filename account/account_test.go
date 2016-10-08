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

//test the SetAccountBalance method and convert it to float (if string)
func TestAccount_AddActivity(t *testing.T) {

	newAccount := Account{}

	result, err := newAccount.AddActivity(
		"26.07.2016",
		"28.07.2016",
		"-116,90",
		"KREDITKARTENABRECHNUNG",
		"VISA-ABR. 555555555555",
		"8888888888",
		"12030088",
		"KARTENZAHLUNG/-ABRECHNUNG",
		"creditor",
		"client",
		"customer")



	if result == false {
		t.Fail()
	}

	if err != nil {
		t.Error(err)
	}

	i, _ := time.Parse("02.01.2006", "26.07.2016")
	if newAccount.Activities[0].AccountingDate != i  {
		t.Log("Fail accounting Date")
		t.Fail()
	}

	i, _ = time.Parse("02.01.2006", "28.07.2016")
	if newAccount.Activities[0].ValueDate != i  {
		t.Log("Fail value Date")
		t.Fail()
	}

	if newAccount.Activities[0].Amount != float64(-116.90)  {
		t.Log("Fail amount Date")
		t.Fail()
	}

	if newAccount.Activities[0].Originator != "KREDITKARTENABRECHNUNG"  {
		t.Fail()
	}

	if newAccount.Activities[0].Reference != "VISA-ABR. 555555555555"  {
		t.Fail()
	}

	if newAccount.Activities[0].AccountNumber != "8888888888"  {
		t.Fail()
	}

	if newAccount.Activities[0].BankCode != "12030088"  {
		t.Log("Fail bankcode")
		t.Fail()
	}

	if newAccount.Activities[0].CreditorId != "creditor"  {
		t.Log("Fail creditor id")
		t.Fail()
	}

	if newAccount.Activities[0].ClientReference != "client"  {
		t.Log("Fail client")
		t.Fail()
	}

	if newAccount.Activities[0].CustomerReference != "customer"  {
		t.Log("Fail customer ref")
		t.Fail()
	}

	if newAccount.Activities[0].Category != "KARTENZAHLUNG/-ABRECHNUNG"  {
		t.Log("Fail category")
		t.Fail()
	}

	if newAccount.Activities[0].Type != -1  {
		t.Log("Fail type")
		t.Fail()
	}

}

func TestAccount_GetActivitiesGroupByMonth(t *testing.T) {

	newAccount := createNewAccountWithActivities()

	activities, err := newAccount.getActivitiesGroupByMonth()

	if len(activities) == 0 {
		t.Fail();
	}

	if err != nil {
		t.Fail()
	}

	if len(activities["07-2016"].Activities) != 2 {
		t.Fail()
	}

	if len(activities["06-2016"].Activities) != 1 {
		t.Fail()
	}

	if activities["07-2016"].Balance != (555.90 - 116.90) {
		t.Fail()
	}

}


func TestAccount_GetActivitiesGroupByMonthUseCache(t *testing.T) {

	newAccount := createNewAccountWithActivities()

	activities, err := newAccount.getActivitiesGroupByMonth()

	if len(activities) == 0 {
		t.Fail();
	}

	if err != nil {
		t.Fail()
	}

	if len(activities["07-2016"].Activities) != 2 {
		t.Fail()
	}

	if len(activities["06-2016"].Activities) != 1 {
		t.Fail()
	}

	activities, err = newAccount.getActivitiesGroupByMonth()

	if len(activities) == 0 {
		t.Fail();
	}

	if err != nil {
		t.Fail()
	}

	if len(activities["07-2016"].Activities) != 2 {
		t.Fail()
	}

	if len(activities["06-2016"].Activities) != 1 {
		t.Fail()
	}

}

//helper function to create a account with multiple activities
func createNewAccountWithActivities() Account {

	newAccount := Account{}

	_, _ = newAccount.AddActivity(
		"26.07.2016",
		"28.07.2016",
		"-116,90",
		"KREDITKARTENABRECHNUNG",
		"VISA-ABR. 555555555555",
		"8888888888",
		"12030088",
		"KARTENZAHLUNG/-ABRECHNUNG",
		"creditor",
		"client",
		"customer")


	_, _ = newAccount.AddActivity(
		"12.07.2016",
		"12.07.2016",
		"555,90",
		"KREDITKARTENABRECHNUNG",
		"VISA-ABR. 555555555555",
		"8888888888",
		"12030088",
		"KARTENZAHLUNG/-ABRECHNUNG",
		"creditor",
		"client",
		"customer")

	_, _ = newAccount.AddActivity(
		"12.06.2016",
		"12.06.2016",
		"222,90",
		"KREDITKARTENABRECHNUNG",
		"VISA-ABR. 2222",
		"8888888888",
		"12030088",
		"KARTENZAHLUNG/-ABRECHNUNG",
		"creditor",
		"client",
		"customer")

	return newAccount
}