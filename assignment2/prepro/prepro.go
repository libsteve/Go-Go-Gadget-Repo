/**
 * A group of functions related to preprocessing a text file
 *
 * Package Usage:
 *	<insert package usage here/>
 */
package prepro

import (
	"os"
	"fmt"
	"buffio"
	"string"
)

/**
 * Start reading a file and processing its contents
 *
 * Parameters:
 *		reader	-	the reader to read the file with
 */
func ReadInput( reader *Reader ) {
	readloop( reader, storedData, make(map [string]string))
}

/**
 * Read the given file and process its contents until a termination keyword is reached
 *
 * Parameters:
 *		reader		-	the reader to read the current file with
 *		storedData	-	the data stored from previous files
 *						i.e. the defined variables
 *		termination -	a map that holds values that cause termination
 */
func readloop( reader *Reader, storedData map [string]string, termination map [string]int ) {
	commands := gencommands( reader, storedData )

	lineString, ok := reader.ReadString(('\n').(byte));
	for ok {
		line, iscommand := getline(reader.ReadString(('\n').(byte)))
		if iscommand {
			if command, ok := getcommand( line ); ok {
				if _, terminate := termination[command]; terminate {
					return
				}
				if function, ok := commands[command]; ok {
					function( remove_hashtag(line) )
					// then print to stdout
				}
			}
		} else {
			insertdefined( line, storedData )
			// then print to stderr
		}
		lineString, ok = reader.ReadString(('\n').(byte));
	}
}

/**
 * Read the given file and until a termination keyword is reached
 *
 * Parameters:
 *		reader		-	the reader to read the current file with
 *		storedData	-	the data stored from previous files
 *						i.e. the defined variables
 *		termination -	a map that holds values that cause termination
 */
func skiploop( reader *Reader, storedData map [string]string, termination map [string]int ) {
	commands := gencommands( reader, storedData )

	lineString, ok := reader.ReadString(('\n').(byte));
	for ok {
		line, iscommand := getline(lineString)
		if iscommand {
			if command, ok := getcommand( line ); ok {
				if _, terminate := termination[command]; terminate {
					return
				}
			}
		}
		lineString, ok = reader.ReadString(('\n').(byte));
	}
}









func gencommands( reader *Reader, storedData map [string]string ) map [string]func( args []string ) {
	commands := make( map [string]func() )

	// add the commands to the map
	commands["#"]		=
		func ( args []string ) {
			os.Stderr("This is technically not a command")
		}, true;

	commands["include"] =
		func ( args []string ) {
			if len(args[1]) > 2 {
				rangemax := len(args[1]) - 1
				filename := args[1]
				filename = filename[1:rangemax]
				if reader, ok := buffio.NewReader( os.Create( filename ) ); ok {
					readloop(newreader, storedData, make(map [string]int))
				} else {
					os.Stderr("Invalid File")
				}
			} else {
				os.Stderr("No File Specified")
			}
		}, true;

	commands["define"]	=
		func ( args []string ) {
			var result string
			defined, _ := args[1]
			for index, word := range args {
				if index > 1 {
					result += word
				}
			}
			storedData[defined] = result
		}, true;

	commands["undef"]	=
		func ( args []string ) {
			if command, ok := storedData[arg[1]]; ok {
				storedData[command] = _, false
			} else {
				os.Stderr("Variable not defined")
			}
		}, true;

	commands["if"]		=
		func ( args []string ) {
			os.Stderr("Warning: This command does nothing now...")
			ifStatement( args[string], reader, storedData )
		}, true;

	commands["ifdef"]	=
		func ( args []string ) {
			ifStatement( args[string], reader, storedData )
		}, true;

	commands["ifndef"]	=
		func ( args []string ) {
			ifStatement( args[string], reader, storedData )
		}, true;

	commands["elseif"]	=
		func ( args []string ) {
			os.Stderr("Warning: This command does nothing now...")
		}, true;

	commands["else"]	=
		func ( args []string ) {
			os.Stderr("Warning: This shouldn't do things by itself...")
		}, true;

	return commands
}







/**
 * get the line from the bitstring
 *
 * take a bitstring and make it a string
 * break it up by spaces into an array of strings
 * return a tupole of an array of strings and a bool
 * the bool is true if the line is a command, false otherwise
 */
func getline( line_string string ) ([]string, bool) {
	line := string.Split( " ", line_string )
	
	iscommand := false
	if command, _ := line[0]; if command == "#" {
		iscommand = true
	} else if command[0] == "#" {
		iscommand = true
	} else {
		// iscommand = false
	}

	return line, iscommand
}

/**
 * get the command from the line
 *
 * takes a line that is already split
 * returns a string and a bool
 * the string is the comand in the line, nil if there is no command
 * the bool is true if the line contains a command. false otherwise
 */
func getcommand( line []string ) (string, bool) {

	if command, ok := line[0]; if command == "#" {
		command, ok = line[1]
	} else if command[0] == "#" {
		command = command[1:len(command)]
	} else {
		os.Stderr("No Command On Line")
		ok = false
		comand = nil
	}

	return command, ok
}

/**
 * remove the hashtag from the beginning of the line
 *
 * takes a line that is already split
 * returns an array string and a bool
 * the array of string is the line without the hashtag, nil if error
 * the bool is true if succes, false otherwise
 */
func remove_hashtag( line []string ) ([]string, bool) {

	if command, ok := line[0]; if command == "#" {
		line = line[1:len(line)]
	} else if command[0] == "#" {
		line[0] = command[1:len(command)]
	} else {
		os.Stderr("No Hashtag On Line")
		ok = false
		comand = nil
	}

	return command, ok
}

/**
 * insert defined words where they are needed
 *
 * takes a line that is already split
 * returns an array of strings and a bool
 * the array of strings is the resulting line with the definitions
 * the bool is true if there was a success, false otherwise
 */
func instertdefined(line []string, storedData map [string]string) ([]string, bool) {
	for index, word := range line {
		if result, ok := storedData[word]; ok {
			line[index] = result
		} else {
			return nil, false
		}
	}
	return line, true
}

// please make sure that line[0] is the command name without the '#'
func ifStatement( args []string, reader *Reader, storedData map [string]string ) {

	conditional := true
	if_type := args[0]
	switch if_type{
	case "ifdef":
		if _, ok := storedData[condition]; ok {
			conditional = true
		} else {
			conditional = false
		}
	case "ifndef":
		if _, ok := storedData[condition]; !ok {
			conditional = true
		} else {
			conditional = false
		}
	case "if":
		// find out if the following statement is true
		conditional = true
	default:
		os.Stderr("How does this sort of thing happen?")
	}
	

	if conditional {
		// continue on loop until else, then skip to the endif
		readloop( reader, storedData, { "else":0; "endif":0; } )
		skiploop( reader, storedData, { "else":0; "endif":0; } )
	} else {
		// skip to the else or endif
		skiploop( reader, storedData, { "else":0; "endif":0; } )
	}
}
