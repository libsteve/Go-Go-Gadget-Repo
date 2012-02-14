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

####Assignment 4 & 5: DAG and make

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

####Assignment 6: Map Reduce

The registry server is a name service to locate services which can be
accessed by remote procedure calls.

The matrix server is a service where matrices can be defined and elements
manipulated.

The clients are interfaces used to connect to the servers.

####Assignment 7: Wagon

A wagon train roams over a plane. The first two wagons appear near the top left and bottom right corner of the plane.

You can move the first wagon up, down, left, or right with the lower case commands u, d, l, or r, respectively, and the last wagon with the corresponding commands in upper case. However, a wagon will only move if it does not leave the plane and if there is no obstacle in the desired direction.

The other wagons will exactly follow the wagon which you moved; e.g., if you move the first wagon, the second wagon will move into the position just vacated by the first wagon, etc.

You can add more wagons to join the train before the first or after the last wagon using the command a in lower or upper case. Each wagon will also appear near the top left and bottom right corner of the plane and will take over the role of first and last wagon, respectively.

####Assignment 8: Games

Implementations of local and distributed versions of the two-player turn-based games Rock Paper Scissors and Tic Tac Toe.

This is a reusable implementation of components based on the Model View Controller paradigm: both games share the code to interact with the user and the network (view), and the local and distributed versions of a game share the code of a referee to enforce the rules (model). Therefore, the interaction between the model and the view must exclusively use this interaction architecture.

Interaction is based on channels. The package netchan implements exporters and importers to distribute channels over a network. This is one way to distribute the games. This paper describes a different pattern to distribute turn-taking games like these over a server which permits only pull access such as remote procedure calls or a protocol such as HTTP.