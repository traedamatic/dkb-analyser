package main

import "fmt"
import (
	"./account"
	"strings"
)

func makeYellow(text string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", text)
}

func PrintHomeScreen(a *account.Account) {


	//clear it
	fmt.Print("\033c")

	//print headline
	fmt.Printf(
		"%s - Analyses your CSV DKB files - Konto: %s - CSV von %s bis %s\033[1E",
		makeYellow("DKB Analyser"),
		a.Title,
		a.BeginDate.Format("02.01.2006"),
		a.EndDate.Format("02.01.2006"))


	fmt.Println(strings.Repeat("-", 200))


}
