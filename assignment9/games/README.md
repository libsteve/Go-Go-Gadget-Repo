Games
=====

by Steven Brunwasser and Kristen Mills


Includes
--------

* main.go
* rps.go
* ttt.go
* local.go
* igame.go
* view.go
* referee.go
* README.md
* Makefile
* test.go
* client.go
* server.go
* remot.go


About
-----

Implementations of local and distributed versions of the two-player turn-based games Rock Paper Scissors and Tic Tac Toe.

This is a reusable implementation of components based on the Model View Controller paradigm: both games share the code to interact with the user and the network (view), and the local and distributed versions of a game share the code of a referee to enforce the rules (model). Therefore, the interaction between the model and the view must exclusively use this interaction architecture.

Interaction is based on channels. The package netchan implements exporters and importers to distribute channels over a network. This is one way to distribute the games. This paper describes a different pattern to distribute turn-taking games like these over a server which permits only pull access such as remote procedure calls or a protocol such as HTTP.


Usage
-----

A makefile has been provided for you to use

To compile the program in the command line, either of
the following will work:
	
	$ make
	$ make all

To remove the executable: 
	
	$ make clean

To generate documentation:

	$ make doc

To run the different programs:
	
	$ ./local [flags] [<filename>] 
	$ ./remote [flags] [port number]
	$ ./server [flags] [port number] channel...
	$ ./test [flags] [host name] channel
	$ ./client [flags] channel

Possible flags for client local and remote are:
	
	-ttt : tic tac toe
	-rps : rock paper scissors

Possible flags for server are:
	
	-port [port number] : specifiy a port number (default is :8080)

Possible flags for test are:

	-host [hostname] : specify host name (default is localhost:8080)

Possible arguments are:

	<filename>	-	the file to use for standard in/out of player 2's view
					good for use with second terminal
					example:

						$ ./games -ttt /dev/ttys002
					
					if no file is given, standard in/out is shared between player 1 and 2

	channel  	-  example:
					
						$ ./server -port :8080 foo
						$ ./test -host localhost:8080 foo
						$ ./test -host localhost:8080 foo=bar
						$ ./server q1 q2
						$ ./client -ttt 1

Notes:
	
	After closing client, if you would like to run it again the server must be restarted. 