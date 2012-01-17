#Homework 4: DAG and make

by Steve Brunwasser and Kristen Mills

##Includes:

	doc.go
	dag.go
	Makefile
	main.go
	dag_implementation.go
	edge_implementation.go


##About:

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

##Usage

A makefile has been provided for you to use

To compile the program in the command line, either of
the following will work:
	
	$ make
	$ make all

To remove the executable: 
	
	$ make clean

To generate documentation:

	$ make doc

To run in the command line:
	
	$./dag [-f file_name] target [more_targets...]

Possible flags:

	-f : specify the mile to read targets and sources from
	without any flags, it should follow the standards of make by 
	reading from either Makefile or makefile
