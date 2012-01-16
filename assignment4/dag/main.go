package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	var force *bool

	force = flag.Bool("f", false, "use the given file as the source")
	flag.Parse()

	var file *os.File
	var error os.Error
	
	/// open the file
	if arg1 := flag.Arg(0); *force {
		if file, error = os.Open(arg1); error != nil {
			fmt.Fprintln(os.Stderr, "404: file not found.")
			return
		}
	} else {
		if file, error = os.Open("Makefile"); error != nil {
			if file, error = os.Open("makefile"); error != nil {
				fmt.Fprintln(os.Stderr, "404: file not found")
				return
			}
		}
	}

	file.Close()

	/// run the command with the file and the rest of the arguments

}
