package main

import (
	"./account"
	"strings"
	"strconv"
	"os/exec"
	"os"
	"fmt"
	"log"
	"sort"
)

var terminalHeight, terminalWidth int

func init()  {
	getTerminalSize()
}

// get the terminal size (width, height)
// @see http://rosettacode.org/wiki/Terminal_control/Dimensions#Go
func getTerminalSize() {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	d, _ := cmd.Output()
	fmt.Sscan(string(d), &terminalHeight, &terminalWidth)
}

//helper function to print a string in yellow
func makeYellow(text string) string {
	return fmt.Sprintf("\033[33m%s\033[39m", text)
}

//helper function to print a string in red
func makeRed(text string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", text)
}

//helper function to print a string in green
func makeGreen(text string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", text)
}

//helper function to print a text
func makeBold(text string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", text)
}

func printHeadline(text string) {
	fmt.Printf("--- \033[1m%s\033[0m %s\033[1E",text, strings.Repeat("-", terminalWidth-len(text)-5))
}

func makeWhitespace(length int) string {
	return strings.Repeat(" ", length)
}

func PrintHomeScreen(a *account.Account) {

	//clear it
	fmt.Print("\033c")

	//print headline
	fmt.Printf(
		"%s - Analyses your CSV DKB files - Konto: %s - CSV from %s until %s\033[1E",
		makeBold(makeYellow("DKB Analyser")),
		a.Title,
		a.BeginDate.Format("02.01.2006"),
		a.EndDate.Format("02.01.2006"))

	//printLine()

	var balanceString string = makeGreen(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	if a.Balance < 0 {
		balanceString = makeRed(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	}

	printHeadline("Totals")

	//print headline
	fmt.Printf(
		"Current balance: %s - Total accounting activities: %s \033[1E",
		makeBold(balanceString),
		strconv.Itoa(len(a.Activities)))

	printHeadline("Months")

	accountingMonths, err := a.GetActivitiesGroupByMonth()

	if err != nil {
		log.Fatal(err)
	}

	//TODO: What happens if month is january?
	lastMonth := int(a.EndDate.Month())
	lastYear := a.EndDate.Year()
	var keys []string = []string{
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth-2), strconv.Itoa(lastYear)),
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth-1), strconv.Itoa(lastYear)),
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth), strconv.Itoa(lastYear))}


	var monthHeadline string

	for _, key := range keys {

		monthHeadline += fmt.Sprintf("%s%s",
			makeWhitespace(terminalWidth/4-7),
			makeBold(accountingMonths[key].Title))
	}

	fmt.Println(monthHeadline)

	var activityLines = make(map[int]string)
	for _, key := range keys {

		for i, activity := range accountingMonths[key].Activities   {

			reference := activity.Reference
			if len(reference) > 20 {
				reference = reference[:20]
			}

			var amountString string = makeGreen(strconv.FormatFloat(activity.Amount, 'f', 2, 64))
			if activity.Amount < 0 {
				amountString = makeRed(strconv.FormatFloat(activity.Amount, 'f', 2, 64))
			}

			var entryLength int = 10+len(reference)+len(strconv.Itoa(i+1))+len(fmt.Sprintf("%.2f", activity.Amount))

			activityLines[i] = activityLines[i] + fmt.Sprintf("#%s - %s - %s - %s%s",
				strconv.Itoa(i+1),
				activity.ValueDate.Format("02.01.2006"),
				reference,
				amountString,
				makeWhitespace(44-entryLength))
		}


	}


	//sorting the activity lines map
	var sortingKey []int
	for k := range activityLines {
		sortingKey = append(sortingKey, k)
	}
	sort.Ints(sortingKey)

	for _, key := range sortingKey {
		fmt.Println(activityLines[key])
	}


}
