package parsearguments

import (
	"errors"
	"regexp"
)

type pArgs struct {
	Filename string
}

// parse the provides os.Args
// returns a error if now file is provided or the file is not a csv file
// returns a instance of the pArgs struct if now error occurs
func ParseArguments(args []string) (pArgs, error) {

	if len(args) < 2 {
		return pArgs{}, errors.New("Wrong amount of cmd args")
	}

	currentArgs := pArgs{Filename: args[1] }

	if matched, _ := regexp.MatchString(".csv", currentArgs.Filename); matched != true {
		return pArgs{}, errors.New("Please provide a csv file")
	}

	return currentArgs, nil
}