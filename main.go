package main

import (
	"os"
	"fmt"
	"./parsearguments"
);

//the main function of the dkb-analyser
func main() {

	pArgs, err := parsearguments.ParseArguments(os.Args)

	if err != nil {
		panic(err)
	}

	//print filename for debug
	fmt.Println(pArgs.Filename)
}
