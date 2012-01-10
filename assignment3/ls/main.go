package main

import (
	"flag"
	"fmt"
	"ls"
)
/*
 * Run the program
 *
 * Usage: 
 *		./ls [args] directory_name
 *		
 * 		possible args:
 *		-R: go through directories recursively
 *		-n: print with information
 *		-t: sort files by modification time
 *
 *		if no arguments are getting, print out alphabetically with 1 file
 *		per line 
 */
func main() {
	var R *bool
	var n *bool
	var t *bool

	R = flag.Bool("R", false, "go through directories recursively")
	n = flag.Bool("n", false, "print with information")
	t = flag.Bool("t", false, "sort files by modification time")
	flag.Parse()

	ls.Ls(flag.Arg(0), *n, *R, *t);

}
