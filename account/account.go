package account

type Account struct {
	Title string
}

// read data from a string array and add it to the account struct
func (a *Account) ReadData(data []string) (bool, error) {

	//check the first entry in the string array
	if data[0] == "Kontonummer:" {

		a.Title = data[1]

		return true, nil

	}

	return false, nil
}