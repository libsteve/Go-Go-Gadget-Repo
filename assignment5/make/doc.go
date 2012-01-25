/*
	A program to create a directed acyclic graph based on target files and 
	their sources. AND THEN MAKE ALL THE THINGS!
	
	Usage:
		./mk [-f file_name] targets
		
		-f :	will force the use of the given file
			
		Will create a DAG based on the given file. AND MAKE ALL THE THINGS
		If no file is given, DAG is produced from either Makefile or makefile. AND MAKE ALL THE THINGS
		If no files are found, an error is displayed and the program stops. AND DON'T MAKE THE THINGS
*/
package documentation

