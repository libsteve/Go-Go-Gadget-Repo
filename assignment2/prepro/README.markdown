#Homework 2: Preprocessor

by Kristen Mills and Steve Brunwasser

##Includes:

	doc.go
	index.html
	pp.go
	prepro.go
	Makefile

##About

A line oriented preprocessor which can read input files and 

* include flies
* (un-)define names
* conditionally ex- or include text if a name is (un-)defined

Handles the following customary commands lines

* process standard input if there are no arguments
* process standard input if an argument is -
* treat any other argument as a file name to be processed
* use a new namespace for each argument

Reports all errors

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

 	$ ./pp -       					  : read from standard in
 	$ ./pp - file1.txt file2.txt ...  : read from standard in as well as files
 	$ ./pp -h						  : launch the webserver
 	$ ./pp file1.txt file2.txt ...    : read from files
 	$ ./pp 							  : read from standard in

Notes:
	
when doing standard in, you should type it in like this

	> line1
	> line2
	> line3
	> line4
	> .

period by itself on a line marks EOF

##Known Bugs

* The webserver prints to standard out.
* Doesn't understand nested conditionals.
* Including files always prints out Invalid File.

