package main

include (
	"os"
	"fmt"
	"flag"
	"bufio"
	"./prepro"
)

func main() {
	if len( os.Args ) < 1 {
		in := bufio.NewReader(os.Stdin)
		/* read from standard in */
	}
	else {
		for _, arg := range args {
			if arg == "-" {
				in := bufio.NewReader(os.Stdin)
				/* read from standard in */
			}
			else {
				in := os.Open(arg)
				/* read from this file*/
				in.Close()
			}
		}	
	}
}
