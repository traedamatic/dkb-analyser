package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
	"./account"
	"strconv"
)

type OverviewLayout struct {
	Header      *gocui.View
	Totals      *gocui.View
	ColumnLeft  *gocui.View
	ColumnMiddle *gocui.View
	ColumnRight *gocui.View
	Footer *gocui.View
	Account *account.Account
	Details *gocui.View
	SelectedActivity string
}

func (l OverviewLayout) Layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	l.Details, _ = g.SetView("details", 0, maxY / 2, maxX - 1, maxY - 1)
	l.Header, _ = g.SetView("header", -1, -1, maxX, 1)
	l.Totals, _ = g.SetView("totals", 0, 2, maxX - 1, 4)
	columnY := int(float32(maxY) - float32(maxY) * float32(0.1))
	l.ColumnLeft, _ = g.SetView("columnLeft", 0, 5, maxX / 3 - 1, columnY)
	l.ColumnMiddle, _ = g.SetView("columnMiddle", maxX / 3, 5, (2) * maxX / 3 - 1, columnY)
	l.ColumnRight, _ = g.SetView("columnRight", 2 * maxX / 3, 5, (3) * maxX / 3 - 1, columnY)
	l.Footer, _ = g.SetView("footer",  0 , columnY + 1, maxX -1 , maxY-1)

	l.ColumnMiddle.Highlight = true
	l.ColumnRight.Highlight = true
	l.ColumnLeft.Highlight = true
	l.ColumnMiddle.SelBgColor = gocui.ColorWhite
	l.ColumnRight.SelBgColor = gocui.ColorWhite
	l.ColumnLeft.SelBgColor = gocui.ColorWhite
	l.ColumnMiddle.SelFgColor = gocui.ColorBlack
	l.ColumnRight.SelFgColor = gocui.ColorBlack
	l.ColumnLeft.SelFgColor = gocui.ColorBlack

	l.Header.Clear()
	l.Totals.Clear()
	l.ColumnLeft.Clear()
	l.ColumnMiddle.Clear()
	l.ColumnRight.Clear()
	l.Footer.Clear()

	writeHeader(l.Header, l.Account)
	writeTotals(l.Totals, l.Account)
	writeFooter(l.Footer)

	lastMonth := int(l.Account.EndDate.Month())
	lastYear := l.Account.EndDate.Year()
	var lastThreeMonthKeys []string = []string{
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth - 2), strconv.Itoa(lastYear)),
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth - 1), strconv.Itoa(lastYear)),
		fmt.Sprintf("0%s-%s", strconv.Itoa(lastMonth), strconv.Itoa(lastYear))}

	for index, key := range lastThreeMonthKeys {

		if index == 0 {
			writeMonth(l.ColumnLeft, l.Account, key)
		}

		if index == 1 {
			writeMonth(l.ColumnMiddle, l.Account, key)
		}

		if index == 2 {
			writeMonth(l.ColumnRight, l.Account, key)
		}

	}

	return nil
}

//helper function to print a string in red
func makeRed(text string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", text)
}

//helper function to print a string in green
func makeGreen(text string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", text)
}

// write little header
func writeHeader(v *gocui.View, a *account.Account) {

	fmt.Fprintf(v, "DKB-Analyser - Analyses your DKB CSV files - Account: %s - CSV from %s until %s", a.Title,
		a.BeginDate.Format("02.01.2006"),
		a.EndDate.Format("02.01.2006"))

}

//write totals
func writeTotals(v *gocui.View, a *account.Account) {

	var balanceString string = makeGreen(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	if a.Balance < 0 {
		balanceString = makeRed(strconv.FormatFloat(a.Balance, 'f', 2, 64))
	}

	fmt.Fprintf(v,
		"Current balance: %s - Total accounting activities: %s ",
		balanceString,
		strconv.Itoa(len(a.Activities)))

}

//write the last month to the terminal ui
func writeMonth(v *gocui.View, a *account.Account, selectedMonth string) {

	v.Title = selectedMonth

	accountingMonths, err := a.GetActivitiesGroupByMonth()

	if err != nil {
		log.Fatal(err)
	}

	for i, activity := range accountingMonths[selectedMonth].Activities   {

		reference := activity.Reference
		if len(reference) > 20 {
			reference = reference[:20]
		}

		var amountString string = makeGreen(strconv.FormatFloat(activity.Amount, 'f', 2, 64))
		if activity.Amount < 0 {
			amountString = makeRed(strconv.FormatFloat(activity.Amount, 'f', 2, 64))
		}

		fmt.Fprintf(v, "#%s - %s - %s - %s\n",
			strconv.Itoa(i+1),
			activity.ValueDate.Format("02.01.2006"),
			reference,
			amountString)
	}

	fmt.Fprint(v, "---\n")

	var monthBalance string = makeGreen(strconv.FormatFloat(accountingMonths[selectedMonth].Balance, 'f', 2, 64))
	if accountingMonths[selectedMonth].Balance < 0 {
		monthBalance = makeRed(strconv.FormatFloat(accountingMonths[selectedMonth].Balance, 'f', 2, 64))
	}

	fmt.Fprintf(v, "Balance: %s", monthBalance)

}

//write footer
func writeFooter(v *gocui.View) {

	fmt.Fprintln(v, "Just a litte side-project to learn the GOLANG...")
	fmt.Fprintln(v, "You can use TAB and arrow keys to navigate through the activities. Press enter to get activities details.")
	fmt.Fprint(v, "For updates and the code visit https://github.com/traedamatic/dkb-analyser")

}

// event handler for tabbing to the next month
func tabToNextMonth(g *gocui.Gui, v *gocui.View) error {

	if v == nil {
		_, err := g.SetCurrentView("columnLeft")
		return err
	}

	switch v.Name() {
	case "columnLeft":
		_, err := g.SetCurrentView("columnMiddle")
		return err
	case "columnRight":
		_, err := g.SetCurrentView("columnLeft")
		return err
	case "columnMiddle":
		_, err := g.SetCurrentView("columnRight")
		return err
	default:
		_, err := g.SetCurrentView("columnLeft")
		return err
	}

	return nil

}

// event handler to move the selection down a line
func moveLineDown(g *gocui.Gui, v *gocui.View) error {

	if v == nil {
		return nil
	}

	v.MoveCursor(0,1,true)

	return nil
}

// event handler to move the selection up a line
func moveLineUp(g *gocui.Gui, v *gocui.View) error {

	if v == nil {
		return nil
	}

	v.MoveCursor(0,-1,true)

	return nil
}

//show the details view
func showDetails(a *account.Account) func(g *gocui.Gui, v *gocui.View) error {
	return func (g *gocui.Gui, v *gocui.View) error {

		if v == nil {
			return nil
		}

		if v.Name() == "details" {
			g.SetViewOnTop("columnRight")
			g.SetViewOnTop("columnMiddle")
			g.SetViewOnTop("columnLeft")
			return nil
		}

		_, row := v.Cursor()
		selectedMonth := v.Title

		detailsView, _ := g.View("details")

		monthAct, _ := a.GetActivitiesGroupByMonth()

		act := monthAct[selectedMonth].Activities[row]

		//TODO: Builds details
		detailsView.Clear()
		fmt.Fprintln(detailsView, "Activity Details")
		fmt.Fprintln(detailsView, "--------------------")
		fmt.Fprintf(detailsView, "Amount: %f\n",act.Amount)
		fmt.Fprintf(detailsView, "ClientReference: %s\n",act.ClientReference)
		fmt.Fprintf(detailsView, "AccountingDate: %s\n",act.AccountingDate)
		fmt.Fprintf(detailsView, "AccountNumber: %s\n",act.AccountNumber)
		fmt.Fprintf(detailsView, "BankCode: %s\n",act.BankCode)
		fmt.Fprintf(detailsView, "Category: %s\n",act.Category)
		fmt.Fprintf(detailsView, "Originator: %s\n",act.Originator)
		fmt.Fprintf(detailsView, "CreditorId: %s\n",act.CreditorId)
		fmt.Fprintf(detailsView, "Reference: %s\n",act.Reference)
		fmt.Fprintf(detailsView, "ValueDate: %s\n",act.ValueDate)
		fmt.Fprintf(detailsView, "Type: %d",act.Type)


		g.SetCurrentView("details")
		g.SetViewOnTop("details")

		return nil
	}
}

func Draw(a *account.Account) {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()


	g.Highlight = true
	g.SelFgColor = gocui.ColorMagenta

	overviewLayout := OverviewLayout{Account: a}

	g.SetManager(overviewLayout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, tabToNextMonth); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, moveLineDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, moveLineUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, showDetails(a)); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}