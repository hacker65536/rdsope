package main

import (
	"os"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	if len(os.Args) == 1 {

		getrds(``)
	} else {
		getrds(os.Args[1])

	}

}
