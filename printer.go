package main

import "fmt"
import (
	"./account"
	"strings"
	"strconv"
)

func makeYellow(text string) string {
	return fmt.Sprintf("\033[33m%s\033[39m", text)
}

func makeRed(text string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", text)
}

func makeGreen(text string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", text)
}


func PrintHomeScreen(a *account.Account) {

	//clear it
	fmt.Print("\033c")

	//print headline
	fmt.Printf(
		"%s - Analyses your CSV DKB files - Konto: %s - CSV from %s until %s\033[1E",
		makeYellow("DKB Analyser"),
		a.Title,
		a.BeginDate.Format("02.01.2006"),
		a.EndDate.Format("02.01.2006"))


	fmt.Println(strings.Repeat("-", 120))

	var balanceString string = makeGreen(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	if a.Balance < 0 {
		balanceString = makeRed(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	}

	//print headline
	fmt.Printf(
		"Current balance %s \033[1E",
		balanceString)

	fmt.Println(strings.Repeat("-", 120))

}
