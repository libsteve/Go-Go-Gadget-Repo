/*
 * A group of functions related to preprocessing a text file
 *
 * Package Usage:
 *	<insert package usage here/>
 */
package pp

include (
	"os"
	"fmt"
)

////////////////////////
////////////////////////
///
/// WARNING!:	The following does not compile.
///				This is just rough pseudo-pseudo-code
///				to get our thoughts together.
///
////////////////////////
////////////////////////

func getCommands() map [string]func(/*place datastructure type here*/) {
	return {
		"define"	: func(/*place datastructure type here*/) {}
		"if"		: ''
		"ifdef"		: ''
		"else"		: ''
		"endif"		: ''
		"include"	: ''
	}
}

func getBuffer(/*input method*/) func() /*buffer return type?*/ {
	
}
