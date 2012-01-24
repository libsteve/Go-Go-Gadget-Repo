/*
A Parser for makefile syntax.
*/
package parser

import (
	"os"
	"bufio"
	"strings"
)

/*
Target Source Command struct

Represents a target, its sources, and the commands associated
*/
type TSC struct {
	Target string
	Sources []string
	Commands []string
}

/*
Make a TSC struct with the given target, sources, and commands
*/
func NewTSC(target string, sources []string, commands []string) *TSC{
	me := new(TSC)
	me.Target = target
	copy(me.Sources, sources)
	copy(me.Commands, commands)
	return me
}

/*
Parse the given file.

Returns:
	an array of pointers to TSC struct instances
	an os.Error if there was a problem reading and interpreting the file
*/
func Parse(file *os.File) []*TSC {

	fileReader := bufio.NewReader(file)

	getResultAndError, readDataFromLine := setUpParser()

	var err os.Error
	err = nil
	for err == nil {
		var line string
		line, err = fileReader.ReadString(byte('\n'))
		if line != "" {
			content_type := findLineContents(line)
			if content_type == 0 {
				readDataFromLine(line)
			}
		}
	}
	return getResultAndError()
}

/*
Sets up the variables for the parser

Returns:
	a function that takes no parameters that returns an array of pointers to TSC structs and an os.Error
	a function that 
*/
func setUpParser() (func() []*TSC, func(string)) {
	var result []*TSC

	var target string
	var sources []string
	var commands []string

	// add the 'current' TSC to the result
	addCurrent := func() {
		if target != "" {
			result = append(result, NewTSC(target, sources, commands))
		}
	}

	resetCurrent := func() {
		target = ""
		sources = make([]string, 0)
		commands = make([]string, 0)
	}

	// allows the parser to return the array of pointers to TSC structs and the os.Error
	getResultAndError := func() []*TSC {
		addCurrent()
		return result
	}

	// allows the parse to read each line and deal with the data
	readDataFromLine := func(line string) {
		if isLineTarget(line) {
			addCurrent()
			resetCurrent()
			var new_command string
			target, sources, new_command = readAsTarget(line)
			commands = append(commands, new_command)
		} else {
			commands = append(commands, readAsCommand(line))
		}
	}

	return getResultAndError, readDataFromLine
}

/*
Read the line as a target

Parameter:
	the line of the target string

Returns:
	a string that represents the target
	an array of strings that reporesent the sources
	a string that represents the first command if there is one
*/
func readAsTarget(line string) (string, []string, string) {
	var target string
	var sources []string
	var command string

	var lineBuffer string
	command_at_end := false

	readSources := func() {
		sources = strings.Split(lineBuffer, " ")
		lineBuffer = ""
	}

	function := map[string]func() {
		":"	:	func() {
			target = lineBuffer
			lineBuffer = ""
		},
		";"	:	func() {
			command_at_end = true
			readSources()
		},
	}

	for _, byte_char := range line {
		char := string(byte_char)
		if funct, ok := function[char]; ok {
			funct()
		} else {
			lineBuffer += char
		}
	}

	if command_at_end {
		command = readAsCommand(lineBuffer)
	} else {
		readSources()
	}

	return target, sources, command
}

/*
Read the line as a command

Parameter:
	a string with a command

Result:
	a string with properly formatted command
*/
func readAsCommand(line string) string {
	return strings.TrimSpace(line)
}

/*
Targets will always start with the first char.
If the line has a non-alphanumeric character first, it is probably a command.

Returns:
	true if the line is a target definition
*/
func isLineTarget(line string) bool {
	if len(line) >= 1 {
		first_char := line[0]
		if trimmed_line := strings.TrimSpace(line); len(trimmed_line) >= 1 {
			new_first_char := trimmed_line[0]
			if first_char == new_first_char {
				// the character is the same when whitespace is removed, thus the line is a target
				return true
			}
		}
	}
	return false
}

/*
Returns:
	-1	-	line is empty
	0	-	line has data (either target:sources...;command or command)
	1	-	line is comment
*/
func findLineContents(line string) int {
	var first_char string

	specialStartChar := map[string]int {
		"#" : 1,
	}

	line_no_spaces := strings.TrimSpace(line)
	if len(line_no_spaces) >= 1 {
		first_char = string(line_no_spaces[0])

		if specialVal, ok := specialStartChar[first_char]; ok {
			// line begins with a special char, return that special value
			return specialVal
		} else {
			// line is either target:sources;command or command
			return 0
		}

	} else {
		// line is empty
		return -1
	}

	// error
	return -10
}












