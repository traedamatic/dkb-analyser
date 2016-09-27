package main

import (
	"os"
	"fmt"
	"./parsearguments"
	"./account"
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

	fmt.Println(pArgs.Filename)

	//start read file
	file, err := os.Open(pArgs.Filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//create account instance
	thisAccount := account.Account{}

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

			_, readErr := thisAccount.ReadData(line)

			if readErr != nil {
				log.Fatal(readErr)
			}


		}

	}

	//debug print out the parsed title
	fmt.Println(thisAccount.Title)


}
