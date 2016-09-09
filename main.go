package main

import (
	"./parsearguments"
	"fmt"
	"os"
	"bufio"
	"strings"
	"encoding/csv"
	"io"
	"log"
);

func main() {

	fmt.Println(parsearguments.ParseArguments())


	//TODO check of args[1] exists
	csvFile := os.Args[1]

	file, err := os.Open(csvFile)

	if err != nil {
		panic(err)
	}

	//TODO: save account data in a struct

	scanner := bufio.NewScanner(file)

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



			if line[0] == "Kontonummer:" {
				fmt.Println("Kontonummer" + line[1])
			}

		}

	}




}
