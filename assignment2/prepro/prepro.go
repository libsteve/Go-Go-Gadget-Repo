/*
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

func getCommands( reader *Reader, storedData map [string]string ) map [string]func( args []string ) {
	var commands map [string]func()

	// add the commands to the map
	commands["#"]		= func ( args []string ) {}
	commands["include"] = func ( args []string ) {}
	commands["define"]	= func ( args []string ) {}
	commands["undef"]	= func ( args []string ) {}
	commands["if"]		= func ( args []string ) {
								ifStatement( args[string], reader, storedData )
							}
	commands["ifdef"]	= func ( args []string ) {
								ifStatement( args[string], reader, storedData )
							}
	commands["ifndef"]	= func ( args []string ) {
								ifStatement( args[string], reader, storedData )
							}
	commands["elseif"]	= func ( args []string ) {}
	commands["else"]	= func ( args []string ) {}

	return commands
}

// takes a line that is already split
func getCommandFromLine( line []string ) string {
	if command, ok := line[0]; !ok {
		os.Stderr("BLEH!!!")
	} else {
		if command == "#" {
			
		}
		else if command[0] == "#" {
			command = command[1:len(command)]
		}
		else {
			os.Stderr("No Command On Line")
		}
	}
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
	
}
