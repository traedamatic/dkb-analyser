package account

/**
 	The Account package holds the account related interfaces and structs. The account struct holds all information
 	of one account.
 */

import (
	"time"
	"strconv"
	"strings"
)

// a generic account interface
type AccountInterface interface {
	SetTitle(string) (bool, error)
	SetBeginDate(string) (bool, error)
	SetEndDate(string) (bool, error)
	SetBalance(string) (bool, error)
	AddActivity(string, string, string, string, string, string, string, ...string) (bool, error)
}

// the account struct holds all account and financial information
type Account struct {
	Title                    string
	BeginDate                time.Time
	EndDate                  time.Time
	Balance                  float64
	Activities               []*Activity

	groupedByMonthActivities map[string]*AccountingMonth
}

// this struct holds all information of one account activity. A activity can be a income or expense.
type Activity struct {
	AccountingDate time.Time
	ValueDate time.Time
	//Buchungstext
	Category string
	Originator string
	Reference string
	AccountNumber string
	BankCode string
	Amount float64
	//Glaeubiger-ID
	CreditorId string
	ClientReference string
	CustomerReference string
	// income or expense (1|-1)
	Type int
}

// this struct holds a stats and the activities of a single month
type AccountingMonth struct {
	//title pattern 01-2006
	Title string
	From time.Time
	Until time.Time
	//next *AccountingMonth
	//prev *AccountingMonth
	Activities []*Activity
	Balance float64
}

// setter method for the title
func (a *Account) SetTitle(title string) (bool, error) {

	a.Title = title
	return true, nil
}

//setter method for the begin date
func (a *Account) SetBeginDate(stringDate string) (bool, error) {

	i, err := time.Parse("02.01.2006", stringDate)

	if err != nil {
		return false, err
	}

	a.BeginDate = i
	return true, nil
}

//setter method for the end date of the csv period
func (a *Account) SetEndDate(stringDate string) (bool, error) {

	i, err := time.Parse("02.01.2006", stringDate)

	if err != nil {
		return false, err
	}

	a.EndDate = i
	return true, nil
}

//setter method for
func (a *Account) SetBalance(newBalance string) (bool, error) {

	floatBalance, err := strconv.ParseFloat(formatCurrency(newBalance), 64)

	if err != nil {
		return false, err
	}

	a.Balance = floatBalance
	return true, nil
}

// formats the german currency format to the english
func formatCurrency(value string) string {

	return strings.Replace(strings.Replace(value, ".", "", 1), ",", ".", 1)
}

// adds a activity to the account struct.
// The additionals should be in order: category, creditor-id, clientReference and customerReference
func (a *Account) AddActivity(  accountingDate,
								valueDate,
								amount,
								originator,
								reference,
								accountNumber,
								bankCode string,
								additionals ...string) (bool, error)  {



	//parse accounting date from string
	aD, err := time.Parse("02.01.2006", accountingDate)

	if err != nil {
		return false, err
	}

	//parse value date from string
	vD, err := time.Parse("02.01.2006", valueDate)

	if err != nil {
		return false, err
	}

	fAmount, err := strconv.ParseFloat(formatCurrency(amount), 64)

	if err != nil {
		return false, err
	}

	var activity *Activity = &Activity{
		AccountingDate: aD,
		ValueDate: vD,
		Amount: fAmount,
		Originator: originator,
		Reference: reference,
		AccountNumber: accountNumber,
		BankCode: bankCode}

	//add additional fields to activity
	for i, additional := range additionals {

		switch i {
		case 0:
			activity.Category = additional
		case 1:
			activity.CreditorId = additional
		case 2:
			activity.ClientReference = additional
		case 3:
			activity.CustomerReference = additional
		}

	}

	//set type
	if activity.Amount < 0 {
		activity.Type = -1
	} else {
		activity.Type = 1
	}

	a.Activities = append(a.Activities, activity)

	return true, nil
}

//return the activities ordered by month as map with string 01-2006 as key
func (a *Account) getActivitiesGroupByMonth() (map[string]*AccountingMonth, error) {

	if len(a.groupedByMonthActivities) > 0 {
		return a.groupedByMonthActivities, nil
	}

	a.groupedByMonthActivities = make(map[string]*AccountingMonth)

	for _, activity := range a.Activities {

		key := activity.ValueDate.Format("01-2006")

		if a.groupedByMonthActivities[key] == nil {

			a.groupedByMonthActivities[key] = &AccountingMonth{
				Balance: activity.Amount,
				Title: key,
				Activities: []*Activity{activity}}


			continue
		}

		a.groupedByMonthActivities[key].Balance += activity.Amount
		a.groupedByMonthActivities[key].Activities = append(a.groupedByMonthActivities[key].Activities, activity)

	}

	return a.groupedByMonthActivities, nil

}
