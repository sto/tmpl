package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	initFlags()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	input := flag.Arg(0)
	if input != "-" {
		inputAbs, err := filepath.Abs(input)
		if err == nil {
			input = inputAbs
		}
		inputsDir = filepath.Dir(inputAbs)
	}

	err := templateRun(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
