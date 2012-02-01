#Homework 6: Map Reduce - Registry and Matrix

by Steve Brunwasser and Kristen Mills

##Includes:

	Makefile
	README.md
	doc-matrix-client.go
	doc-matrix-server.go
	doc-registry-client.go
	doc-registry-server.go
	imatrix.go
	iregistry.go
	matrixclient.go
	matrixcommon.go
	matrixserver.6
	matrixserver.go
	parser.go
	reducer.go
	regclient-run.go
	regclient.go
	registry.go
	regserver.go


##About:

The registry server is a name service to locate services which can be
accessed by remote procedure calls.

The matrix server is a service where matrices can be defined and elements
manipulated.

The clients are interfaces used to connect to the servers.

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

To run the matrix server in the command line:
	
	$./matrixServer

To run the matrix client in the command line:
	
	$./matrixClient

To run the registry server in the command line:

	$./regserver

To run the registry client in the command line:

	$./regclient

The servers run until they are exited with Ctrl-C.

The matrix client runs until given the following command:

	Close

The registry client runs until given the following command:

	close

The following are commands for the registry client with user input:

	Command:				Description:

	help					display a list of commands
	bind: KEY, VALUE 		bind the KEY to the VALUE
	lookup: KEY				find the VALUE for th KEY
	remove: KEY				remove the KEY and the VALUE
	close 					close the session

The following are commands for the matrix client with user input:

	Command:				Description:
	Make:NAME,ROWS,COLS		Create matrix NAME with ROWS rows and COLS cols
	Dim:NAME				Returns the dimensions of matrix NAME
	Remove:NAME				Remove matrix NAME from the database
	Get:NAME,I,J			Get the value in matrix NAME at I,J
	SET:NAME,I,J,VALUE		Set the value in matrix NAME at I,J to VALUE
	Close					Disconnect from the server and close the program


	
