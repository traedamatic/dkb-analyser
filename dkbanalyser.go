package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"io"
	"./parsearguments"
	"./csvdata"
	"./account"
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

	////debug print out the parsed title
	//fmt.Println(thisAccount.Title)
	//fmt.Println(thisAccount.BeginDate.String())
	//fmt.Println(thisAccount.EndDate.String())
	//fmt.Println(len(thisAccount.Activities))


	PrintHomeScreen(thisAccount)


	// this scans the stdin
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}