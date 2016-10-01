package account

/**
 	The Account package holds the account related interfaces and structs. The account struct holds all information
 	of one account.
 */

import (
	"time"
)

// a generic account interface
type AccountInterface interface {
	SetTitle(string) (bool, error)
	SetBeginDate(string) (bool, error)
}

/// the account struct holds all account and financial information
type Account struct {
	Title string
	BeginDate time.Time
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