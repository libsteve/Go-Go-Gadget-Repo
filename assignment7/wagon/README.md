Wagon Train
===========

by Steven Brunwasser and Kristen Mills


Includes
--------

* Makefile
* main.go
* wagon.go
* screen.go
* README.md


About
-----

A wagon train roams over a plane. The first two wagons appear near the top left and bottom right corner of the plane.

You can move the first wagon up, down, left, or right with the lower case commands u, d, l, or r, respectively, and the last wagon with the corresponding commands in upper case. However, a wagon will only move if it does not leave the plane and if there is no obstacle in the desired direction.

The other wagons will exactly follow the wagon which you moved; e.g., if you move the first wagon, the second wagon will move into the position just vacated by the first wagon, etc.

You can add more wagons to join the train before the first or after the last wagon using the command a in lower or upper case. Each wagon will also appear near the top left and bottom right corner of the plane and will take over the role of first and last wagon, respectively.


Instructions
------------

To build and compile:

	$ make

To run the program:

	$ ./wag

To move the head of the wagon, press:

	u 	-	move up
	d 	-	move down
	l 	-	move left
	r 	-	move right

	up arrow 	-	move up
	down arrow 	-	move down
	left arrow 	-	move left
	right arrow	-	move right

To move the tail of the wagon, press:

	U 	-	move up
	D 	-	move down
	L 	-	move left
	R 	-	move right

To add a wheel to the wagon train, press:

	a 	-	add a wheel to the top of the screen and set it as the wagon's head
	A 	-	add a wheel to the bottom of the screen and set it as the wagon's tail

To quit the program, press:

	q 		-	quit the program
	ESC-ESC	-	quit the program