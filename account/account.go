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
}

/// the account struct holds all account and financial information
type Account struct {
	Title string
	BeginDate time.Time
	EndDate time.Time
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
