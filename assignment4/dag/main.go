package main

import (
	"flag"
	"os"
	"fmt"
	"./dag"
	"bufio"
	"strings"
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

	parsedLines := parser(file)

	thedag := dag.MakeDag()

	for _, target := range parsedLines {
		thedag.Add([]string{target[0]}, target[1:], dag.MakeEdge())
	}
	for i, arg := range flag.Args(){
		if i != 0 {
			thedag.Apply(arg)
		}
	}
}

/**
 * Reads the specified file and parses its contents.
 *
 * Parameters:
 *		file *os.File - a pointer to the file
 *
 * Returns:
 *		A 2D array of strings where the first entry of the 
 *		second-dimensional array is the target and the following 
 *		entries are the sources related to the target
 *
 * Known Bugs:
 *		If there is a tab immediately between the colin and the first source, 
 *		then that source will not be added to the resulting list.
 */
func parser(file *os.File) [][]string {
	result := make([][]string, 0)

	fileReader := bufio.NewReader(file)
	line, err := fileReader.ReadString(byte('\n'))
	for err == nil {
		parse(line, &result)
		line, err = fileReader.ReadString(byte('\n'))
	}
	parse(line, &result)

	return result
}

func parse(line string, result *[][]string){
	targetResult := strings.Split(line[:len(line)-1], ":")
	if(len(targetResult) > 1){
		sourcesCommands := strings.Split(targetResult[1], ";")
		if len(sourcesCommands) > 0{
			sources :=strings.Split(sourcesCommands[0], " ")
			target := []string{targetResult[0]}
			for _, source := range sources[1:]{
				target = append(target, strings.TrimSpace(source))
			}
			//deal with commands next week
			*result = append(*result, target)
		}else{
			//deal with commands next week
		}
	}
}
