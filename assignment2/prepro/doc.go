/*
A simple preprocessor that takes in input files and can
  
  		include files,
  		(un-)define names,
 		and conditionally ex- or include text if a name is (un-)defined.
 
Usage:
  		./pp -       					: read from standard in
 		./pp - file1.txt file2.txt ...  : read from standard in as well as files
 		./pp -h							: launch the webserver
 		./pp file1.txt file2.txt ...	: read from files
 		./pp 							: read from standard in 
 */
package documentation
