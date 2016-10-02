package main

import (
	"os"
	"fmt"
	"./parsearguments"
	"./account"
	"./csvdata"
	"log"
	"bufio"
	"encoding/csv"
	"strings"
	"io"
);

//the main function of the dkb-analyser
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


	printHomeScreen(thisAccount)


	// this scans the stdin
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}


}

func makeYellow(text string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", text)
}

func printHomeScreen(a *account.Account) {

	//clear it
	fmt.Print("\033c")
	fmt.Printf("%s - Analyses your CSV DKB files - Konto: %s - CSV von %s bis %s", makeYellow("DKB Analyser"), a.Title, a.BeginDate.Format("01.02.2006"), a.EndDate.Format("01.02.2006"))


}
