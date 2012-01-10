/**
 * A group of functions related to listing all files in a given directory.
 *
 * ls.Ls() takes 2 parameters:
 *		1) the directory name
 *		2) a list of parameters
 *
 * Parameters can be any of the following:
 *		none - list one line at a time, display files in alphabetical order, no recursive directory searching
 *		"-n" - display with information
 *		"-R" - go through directories recursively
 *		"-t" - sort by timestamp
 *
 * Usage:
 *		ls.Ls(directoryname, argumentlist)
 */
package ls

import (
	"io"
	"io/ioutil"
	"os"
	"sort"
	"fmt"
)

func Ls(dirname string, n bool, R bool, t bool) string {
	// args can be any of the following:
	//		default
	//			list one line at a time display files in alphabetical order
	//		-n
	//			display with information
	//		-R
	//			go through directories recursively
	//		-t
	//			sort by timestamp
	lsdir := readdir
	sort := alphasort
	disp := namedisp

	if n {
		disp = infodisp
	}
	if R {
		lsdir = recurdir
	}
	if t {
		sort = timesort
	}

	result := ""
	files := lsdir(dirname, sort, disp)
	for index, file := range files {
		result += file + "\n"
	}
	return result
}

// function to go through directories recursively
func recurdir(filename string, sort func(files []*os.FileInfo) []*os.FileInfo, disp func(*os.FileInfo) string) ([]string, os.Err) {
	// open the file
	// if the file is a folder, for each file in it
	//		use sort() to sort the files
	//		for each file
	//			if the file is a folder, resurse and save in a queue for later use
	// use result from sort() and put it through disp() to get the display strings
	// for each file in the queue
	//		put it through disp() to get display string and add it to the end of the string array
	// return the display strings
	display := make(map[int]string)
	if fi, ok = os.Stat(filename); ok {
		if fi.IsDirectory() {
			files := ioutil.ReadDir(filename)
			files = sort(files)
			for index, file := range files {
				if file.IsDirectory() {
					morefiles := recurdir(file.Name, sort, disp)
					// add morefiles to queue
				}
				display[index] = disp(file), true
			}
		}
	} else {
		return ok
	}
}
// function not go recursively through directories
func readdir(filename string, sort func(files []*os.FileInfo) []*os.FileInfo, disp func(*os.FileInfo) string) []string {
	// open the file
	// if the file is a folder, for each file in it
	//		use sort() to sort the files
	// use result from sort() and put it through disp() to get the display strings
	// return the display strings
}

// sort function to sort alphabetically
func alphasort(files []*os.FileInfo) []*os.FileInfo {
	// use sort.Interface somehow to sort the files
}
// sort function to sort by timestamp
func timesort(files []*os.FileInfo) []*os.FileInfo {
	// use sort.Interface somehow to sort the files
}

// display function to only display the name
func namedisp(file *os.FileInfo) string {
	return file.Name + "\n"
}
// display function to display all information
func infodisp(file *os.FileInfo) string {
	file.Mode
	file.Blocks
	file.Blksize

	file.size
}

