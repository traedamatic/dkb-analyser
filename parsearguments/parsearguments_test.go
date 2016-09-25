package parsearguments

import (
	"testing"
)

// test for no second argument
func TestParseArgumentsForMissingFile(t *testing.T) {

	var args = []string{"main.go"}

	_, err := ParseArguments(args)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "Wrong amount of cmd args" {
		t.Fail()
	}
}

// test is second argument will be set as Filename in pArgs struct
func TestParseArgumentsWithSecondArg(t *testing.T) {

	var args  = []string{"main.go", "/path/to/tmp.csv"}

	pArgs, err := ParseArguments(args)

	if err != nil {
		t.Fail()
	}

	if	pArgs.Filename == "" {
		t.Fail()
	}

	if pArgs.Filename != "/path/to/tmp.csv" {
		t.Fail()
	}
}

//test if the file has the csv extension
func TestParseArgumentsFilenameWithoutCsv(t *testing.T) {

	var args  = []string{"main.go", "/path/to/tmp.png"}

	_, err := ParseArguments(args)

	if err == nil {
		t.Fail()
	}

	if err.Error() != "Please provide a csv file" {
		t.Fail()
	}

}
