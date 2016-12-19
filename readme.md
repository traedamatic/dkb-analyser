# DKB Analyser

**Motivation:** I want to learn GO. So I decided to write a little DKB CSV file 
Analyser. My personal goal is to learn the basic concepts of GO in a nice
and relaxed way. Comments and feedback is always weclome

**Goal:** Parse a DKB CSV file and print out a nice report of the bank
account details. First it will be command-line tool only.

# Main Tasks

* Write tests for all code.
* Parse command line args and check if csv file isset and exists
* Parse Header from csv custom dkb header
* Print out the header
* Calculate the income and expenses
* Forecast calculation for the next month

# Dependencies

* github.com/jroimartin/gocui / fantastic termbox ui implementation

# Usage

```
go run main.go CSVFILEPATH
```

# Tests

This project has multiple little packages inside. All package will have tests.

Within every package you can run:

```
go test 
```

# License

Nicolas Traeder / MIT