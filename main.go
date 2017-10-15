package main

import (
	"log"
	"bufio"
	"os"
	"encoding/csv"
	"strings"
	"io"
	"./parsearguments"
	"./csvdata"
	"./account"
	"./ui"
)

func main() {

	pArgs, err := parsearguments.ParseArguments(os.Args)

	if err != nil {
		panic(err)
	}

	//start read file
	file, err := os.Open(pArgs.Filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//create account instance
	thisAccount := &account.Account{}

	for scanner.Scan() {

		r := csv.NewReader(strings.NewReader(scanner.Text()))
		r.Comma = ';'

		for  {

			line, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			_, readErr := csvdata.ParseCSV(thisAccount, line)

			if readErr != nil {
				log.Fatal(readErr)
			}


		}

	}

	// start the ui
	ui.Draw(thisAccount)

}