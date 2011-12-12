package main

include (
	"os"
	"fmt"
	"./prepro"
)

func main() {
	if len( os.Args ) < 1 {
		/* read from standard in */
	}
	else {
		for _, arg := range args {
			if arg == "-" {
				/* read from standard in */
			}
			else {
				/* read from this file*/
			}
		}	
	}
}
