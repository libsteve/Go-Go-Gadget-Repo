/*
A Simple Parser for Commands and Input

Usage:
	get a Commands struct with NewCommands()
	add commands with AddCommand(keyword, function)
	add commands that take input with AddInputCommand(keyword, function)
	parse one line of commands and execute the command with Parseln(line)
	
The default command-argument separator token is ":"
	you can change it by changing the value of the ComSep variable in the Commands struct

The default argument separator token is ","
	you can change it by changing the value of the ParamSep variable in the Commands struct

Whitespace between arguments is removed, but the whitespace within arguments is preserved
	"command:      all_this_whitespace_is_removed          ,           same_here      "
	"command:this     whitespace      is       not     removed,neither    is      this"

Commands are whitespace sensitive.
	"These     are    different      commands"
	"These are different commands            "
	"            These are different commands"
*/
package parser

import (
	"os"
	"strings"
)

/*
This is the type that functions must conform to if they are to be commands
*/
type Com func()os.Error

/*
This is the type that functions must conform to if they are to be commands with arguments
*/
type Incom func([]string)os.Error

/*
The Commands Struct
*/
type Commands struct {
	Coms map[string]Com
	Incoms map[string]Incom
	ComSep string
	ParamSep string
}

/*
Get a new Commands struct
*/
func NewCommands() *Commands {
	commands := new(Commands)
	commands.Coms = make(map[string]Com)
	commands.Incoms = make(map[string]Incom)

	commands.ComSep = ":"
	commands.ParamSep = ","
	return commands
}

/*
Add a command to the Commands struct

Usage:
	AddCommand(keyword, function)

Parameters:
	name - the string to signal the use of the command when parsed
	command - the function to be used when the command is executed

The command/function will return either nil or os.Error
See the type definition for Com
*/
func (c *Commands) AddCommand(name string, command Com) {
	c.Coms[name] = command
}

/*
Add a command that takes arguments to the Commands struct

Usage:
	AddCommand(keyword, function)

Parameters:
	name - the string to signal the use of the command when parsed
	command - the function to be used when the command is executed

The command/function will be passed an array of string arguments to handle as the arguments
The command/function will return either nil or os.Error
See the type definition for Incom
*/
func (c *Commands) AddInputCommand(name string, command Incom) {
	c.Incoms[name] = command
}

/*
separate the args string by the Commands.ParamSep token and return the resulting array of strings
*/
func (c *Commands) parseArgs(args string) []string {
	tokens := strings.Split(args, c.ParamSep)
	for index, token := range tokens {
		tokens[index] = strings.TrimSpace(token)
	}
	return tokens
}

/*
Parse the string into a command and it's arguments if any

Usage:
	Parseln(line)

Parameters:
	line - the line to read one command and it's possible arguments from

Returns:
	os.Error
	- the command/function's os.Error is returned if the line was correctly parsed
	- if multiple command tokens are found on the line, the "Too Many Command Separators" error is returned
	- if an argument-accepting command was not found, the "Command:[COMMAND:ARGS] Not Found" error is returned
	- if a command was not found, the "Command:[COMMAND] Not Found" error is returned
*/
func (c *Commands) Parseln(line string) os.Error {
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, c.ComSep)
	length := len(tokens)
	switch length {
		case 1:
			return c.runCom(tokens[0])
		case 2:
			if tokens[1] == "" {
				return c.runCom(tokens[0])
			}
			return c.runIncom(tokens[0], tokens[1])
		default:
			return os.NewError("Too Many Command Separators")
	}
	return os.NewError("Something Failed")
}

/*
run the supplied command with the following name
*/
func (c *Commands) runCom(token string) os.Error {
	com, ok := c.Coms[token]
	if ok {
		return com()
	}
	return os.NewError("Command:["+token+"] Not Found")
}

/*
run the supplied input-command with the following name and raw-args string
*/
func (c *Commands) runIncom(token string, args_str string) os.Error {
	args := c.parseArgs(args_str)
	com, ok := c.Incoms[token]
	if ok {
		return com(args)
	}
	return os.NewError("Command:["+token+":ARGS] Not Found")
}

/*
Tests to make sure the package works

Prints to screen the following message if all goes well:
	Test is successful
	Test is successful
	Facebook Friends:
	~Steve Brunwasser
	~Kristen Mills
	Command:[will fail] Not Found
	Command:[test:ARGS] Not Found
	Command:[fb] Not Found <-no arguments
*/
func Test() {
	com := NewCommands()
	f := func() os.Error { 
		println("Test is successful")
		return nil
	}
	fb := func(args []string) os.Error { 
		println("Facebook Friends:")
		for _, t := range args {
			println("~"+t)
		}
		return nil
	}
	com.AddCommand("test", f)
	com.AddInputCommand("fb", fb)
	com.Parseln("test")
	com.Parseln("test:")
	com.Parseln("fb: Steve Brunwasser, Kristen Mills")
	println(com.Parseln("will fail").String())
	println(com.Parseln("test: will fail").String())
	println(com.Parseln("fb:").String() + " <-no arguments")

}