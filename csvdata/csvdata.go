package csvdata

import (
	"../account"
)

/**
	The csv-data package will parse a CSV slice and write the parsed data to an account.
 */

// read data from a string array and add it to the account struct
func ParseCSV(account account.AccountInterface, data []string) (bool, error) {

	if len(data) == 0 {
		return false, nil
	}

	//check the first entry in the string array
	switch data[0] {

	case "Kontonummer:" :
		return account.SetTitle(data[1])

	case "Von:":
		return account.SetBeginDate(data[1])
	}

	return false, nil
}