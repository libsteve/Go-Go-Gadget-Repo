package main
import (
	"flag"
	"os"
	"fmt"
	"./dag"
	"./edge_implementation"
	"./parser"
	"bufio"
	"strings"
)

func main() {
	var force *string

	force = flag.String("f", "Makefile", "use the given file as the source")
	flag.Parse()

	var file *os.File
	var error os.Error

	/// open the file
	if file, error = os.Open(*force); error != nil {
		fmt.Fprintln(os.Stderr, "404: file not found.")
		return
	}
	parsedLines = parser(file)

	thedag :=  dag.New(20)

	for _, tsc := range parsedLines {
		thedag.Add(tsc->Target, tsc->Sources, *dag.MakeEdge(tsc->Commands))
	}
	for i, arg := range flag.Args(){
		if i != 0 {
			fmt.Println(arg + ":")
			thedag.Apply(arg)
			if i != flag.NArg()-1 {
				fmt.Println()
			}
		}
	}
}
