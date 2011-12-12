/**
 * A group of functions related to preprocessing a text file
 *
 * Package Usage:
 *	<insert package usage here/>
 */
package prepro

include (
	"os"
	"fmt"
	"buffio"
	"string"
)

////////////////////////
////////////////////////
///
/// WARNING!:	I have no real idea what I am doing...
///
////////////////////////
////////////////////////

/**
 * Start reading a file and processing its contents
 *
 * Parameters:
 *		reader	-	the reader to read the file with
 */
func Runloop( reader *Reader ) {
	runloop( reader, var storedData map [string]string )
}

/**
 * Read the given file and process its contents
 *
 * Parameters:
 *		reader		-	the reader to read the current file with
 *		storedData	-	the data stored from previous files
 *						i.e. the defined variables
 */
func runloop( reader *Reader, storedData map [string]string ) {
	commands := gencommands( reader, storedData )

	for /*the amount of lines in the file*/ {
		line, iscommand := getline(reader.Readline())
		if iscommand {
			if command, ok := getcommand( line ); ok {
				if function, ok := commands[command]; ok {
					function( remove_hashtag(line) )
				}
			}
		} else {
			// scan for defined variables
			// replace defined variables
		}
	}
}






/**
 ******************************
 ******************************
 ** THIS IS HORRIBLE, MUST FIX
 ******************************
 ******************************
 */
func gencommands( reader *Reader, storedData map [string]string ) map [string]func( args []string ) {
	var commands map [string]func()

	// add the commands to the map
	commands["#"]		= func ( args []string ) {}
	commands["include"] = func ( args []string ) {}
	commands["define"]	= func ( args []string ) {
		
	}
	commands["undef"]	= func ( args []string ) {}

	////////////////
	// if statements
	commands["if"]		= func ( args []string ) {
		ifStatement( args[string], reader, storedData )
	}
	commands["ifdef"]	= func ( args []string ) {
		ifStatement( args[string], reader, storedData )
	}
	commands["ifndef"]	= func ( args []string ) {
		ifStatement( args[string], reader, storedData )
	}
	////////////////

	commands["elseif"]	= func ( args []string ) {}
	commands["else"]	= func ( args []string ) {}

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
func getline( line_bitstring []byte ) ([]string, bool) {
	line := string.Split( " ", line_bitstring.(string))
	
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
 ******************************
 ******************************
 ** THIS IS HORRIBLE, MUST FIX
 ******************************
 ******************************
 */
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
	
}
