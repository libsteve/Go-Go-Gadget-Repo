/*
	A program to create a directed acyclic graph based on target files and 
	their sources.
	
	Usage:
		./dag [-f file_name] targets
		
		-f :	will force the use of the given file
			
		Will create a DAG based on the given file.
		If no file is given, DAG is produced from either Makefile or makefile.
		If no files are found, an error is displayed and the program stops.
*/
package documentation

