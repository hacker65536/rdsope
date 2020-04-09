package main

import (
	"os"
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
