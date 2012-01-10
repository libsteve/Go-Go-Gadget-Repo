package main

import (
	//"os"
	"flag"
	"fmt"
	"ls"
)
/*
 *
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
