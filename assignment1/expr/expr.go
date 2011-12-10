/*
Finds the resulting value from a given expression

Commandline Usage:
	expr expression
*/
package main

import (
  "fmt"
  "os"
  "strconv"
)

/*
Run the program.

Uses os.Args arguments.
os.Args should be something like this:
	[expr 3 + ( ( 1 + 2 ) * ( 2 + 3 ) ) / ( 2 - 1 )]
*/
func main() {
	if len( os.Args ) > 1 {
		fmt.Println( strconv.Itoa( Evaluate( os.Args ) ) )
	} else {
		fmt.Println(os.Args[0] + ": syntax error")
		os.Exit(2)
	}
}
