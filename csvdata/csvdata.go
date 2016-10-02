package csvdata

import (
	"../account"
	"regexp"
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
	case "Bis:":
		return account.SetEndDate(data[1])
	case "Kontostand vom:":
		return account.SetBalance(data[1])
	case "Buchungstag":

		//The row is the headers for the activities
		//Do we need the header in some way?
		return true, nil
	}

	//from the all entries should be account activities
	// but check for date with regxep and skip if no match
	regexpPattern := "^[0-3]{1}[0-9]{1}.(01|02|03|04|05|06|07|08|09|10|11|12).20[0-9]{2}$"
	if matched, err := regexp.MatchString(regexpPattern, data[0]); matched == false || err != nil {

		return false, nil
	}

	return account.AddActivity(
		data[0],
		data[1],
		data[7],
		data[3],
		data[4],
		data[5],
		data[6],
		data[2],
		data[8],
		data[9],
		data[10],
	)
}