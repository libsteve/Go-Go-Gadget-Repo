package main

import (
	"flag"
	"os"
)

func main() {
	var force *bool

	force = flag.Bool("-f", false, "use the given file as the source")
	flag.Parse()

	args := flag.Args()

	var file os.File
	
	/// open the file
	if ok, arg1 := agrs[1]; ok == nil && force {
		if file, error := os.Open(arg1); error != nil {
			os.Stderr("404: file not found.")
			return
	} else {
		if file, error := os.Open("Makefile"); error != nil {
			if file, error = os.Open("makefile"); error != nil {
				os.Stderr("404:file not found")
				return
			}
		}
	}

	/// run the command with the file and the rest of the arguments

}
