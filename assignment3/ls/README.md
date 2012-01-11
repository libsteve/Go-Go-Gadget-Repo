#Homework 3: Templates

by Kristen Mills and Steve Brunwasser

##Includes:

	doc.go
	ls.go
	Makefile
	main.go


##About:

An implentation of a simplified version of the unix command ls

Handles the following command lines

* Traversal and formatting for option -n
* Template for option -R
* Support for option -t

##Usage

A makefile has been provided for you to use

To compile the program in the command line, either of
the following will work:
	
	$ make
	$ make all

To run in the command line:
	
	$./pp [flags] [files]

To remove the executable: 
	
	$ make clean

To generate documentation:

	$ make doc

Possible running commands:

	$ ./ls [flags] directory

Possible flags:

	-n : print out file information
	-R : Recursively go through directories
	-t : sort by timestamp
	without any flags, it should follow the standards of ls -1 with one 		file per line sorted alphabetically 
