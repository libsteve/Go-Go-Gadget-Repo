#Homework 1: Arithmetic 

by Kristen Mills and Steve Brunwasser

##Includes:
	expr.go
 	expreval.go
	Makefile

##About

Functions the same way as the unix command expr but 
is restricted to integer
arithmetic.

##Usage

We have provided a Makefile for you to use.

To compile the program in the command line, either of 
the following will work:
	$ make
	$ make all

To run in command line:
	$ ./expr (expression)

To generate documentation:
	$ make doc

To run the tests in the command line:
	$ make test
	
	Note: For each test, it is in the form 
	./expr 1 + 2
	3
	expr 1 + 2
	3 

To remove the executable:
	$ make clean

Example use:

	$ make
	6g  -o _go_.6 expreval.go expr.go 
	6l  -o expr _go_.6
	$ ./expr 2 \* \( 3 + 4 \) / \( 2 - 4 \)
	-7
