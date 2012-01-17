#Go Go Gadget 

by Kristen Mills and Steve Brunwasser

##About

These are some the assignments we are doing for one of our classes, Systems and Concurrent Programming in Go.

Links

[Go Website](http://golang.org/)

##Includes

####Assignment 1: Arithmetic
	
An implementation of the unix command expr restricted to integer division

####Assignment 2: Preprocessor

A line oriented preprocessor which can read input files and 

* include flies
* (un-)define names
* conditionally ex- or include text if a name is (un-)defined

Handles the following customary commands lines

* process standard input if there are no arguments
* process standard input if an argument is -
* treat any other argument as a file name to be processed
* use a new namespace for each argument

Reports all errors.

####Assignment 3: Templates

An implentation of a simplified version of the unix command ls

Handles the following command lines

* Traversal and formatting for option -n
* Template for option -R
* Support for option -t

####Assignment 4: DAG and make

The DAG traversal command loads a mkfile - which need not include commands -
and for the first target, or for each target name on the command line, 
displays the target lines depth-first, i.e., prerequisites before targets

A make-like command loads a mkfile — which would include commands — and for
the first target, or for all target names on the command line, creates a 
shell script to execute commands if targets do not exist or prerequisites 
are newer.

* A comment extends from a special character (typically #) to the end of the
line and is ignored.
* Completely blank lines are ignored.
* A dependency (one or more edges of the DAG) is a block of one or more 
lines.
* The first line of such a block is not indented. It contains one or more 
target names, zero or more source names, and optional text; targets have 
the sources as prerequisites. Targets and sources are separated by a special
character (typically :). The optional text is preceded by a special 
character (typically ;). Names should not contain special characters and 
are separated from each other by white space.
* Further lines in the block must be indented and are part of the optional 
text. For the purposes of make, the optional text consists of shell commands
which are expected to (re-) create the targets of the dependency.
