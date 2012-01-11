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
	"ntainer/vector"
	"sort"
	"fmt"
)

type FileData struct{
	Mode string
	Nlink uint64
	Uid int
	Gid int
	Size int64
	Mtime string
	name string
}

func Ls(dirname string, n bool, R bool, t bool) []FileData {
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

//	if n {
//  	disp = infodisp
//	}
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
func recurdir(filename string, sort func(files []*os.FileInfo) []*os.FileInfo, disp func(*os.FileInfo) string) (vector.StringVector, os.Err) {
	// open the file
	// if the file is a folder, for each file in it
	//		use sort() to sort the files
	//		for each file
	//			if the file is a folder, resurse and save in a queue for later use
	// use result from sort() and put it through disp() to get the display strings
	// for each file in the queue
	//		put it through disp() to get display string and add it to the end of the string array
	// return the display strings

	displays := new(vector.StringVector)
	dirqueue := new(vector.StringVector)
	if fi, ok = os.Stat(filename); ok {
		if fi.IsDirectory() {
			files := ioutil.ReadDir(filename)
			files = sort(files)
			for index, file := range files {
				if file.IsDirectory() {
					morefiles := recurdir(file.Name, sort, disp)
					dirqueue.Push("")
					dirqueue.AppendVector(morefiles)
				}
				displays.Push(disp(file))
			}
		}
	for index, fidisp := range *dirqueue {
			displays.Push(fidisp)
		}
		return displays, ok
	} else {
		return displays, ok
	}
}
// function not go recursively through directories
func readdir(filename string, sort func(files []*os.FileInfo) []*os.FileInfo, disp func(*os.FileInfo) string) (vector.StringVector, os.Err) {
	// open the file
	// if the file is a folder, for each file in it
	//		use sort() to sort the files
	// use result from sort() and put it through disp() to get the display strings
	// return the display strings

	displays := new(vector.StringVector)
	if fi, ok = os.Stat(filename); ok {
		if fi.IsDirectory() {
			files := ioutil.ReadDir(filename)
			files = sort(files)
			for index, file := range files {
				displays.Push(disp(file))
			}
		}
		return displays, ok
	} else {
		return displays, ok
	}

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
	return file.Name
}
// display function to display all information
func infodisp(file *os.FileInfo) template {
//	t := time.NanosecondsToLocalTime(file.Mtime_ns)
//	timeStr := t.Format("Jan _2 15:04")
	node := fileInfoToNode(file)
	temp := template.Must(template.New("ls").Parse("{{.Mode}}  {{.Nlink}}  {{.Uid}}  {{.Gid}}  {{printf `%7d` .Size}} {{.Mtime_ns}}  {{.Name}}\n"))
	//str := fmt.Sprintf("%d  %d  %d  %d  %8d  %s  %s\n", file.Mode,
	//file.Nlink,	file.Uid, file.Gid, file.Size, timeStr, file.Name)
	return temp
}

func fileInfoToNode(file *os.FileInfo) Node{
	t := time.NanosecondsToLocalTime(file.Mtime_ns);
	timeStr := t.Format("Jan _2 15:04");
	permissions :=""
	permo := fmt.Sprintf("%o", file1.Mode)
	rwx := permo[len(permo)-3: len(permo)]
	if permo[0] == '4'{
		permissions+="d"
	}else{
		permissions += "-"
	}

	for _, char := range rwx {
		switch char {
		case '0':
			permissions+="---"
		case '1':
			permissions+="--x"
		case '2':
			permissions+="-w-"
		case '3':
			permissions+="-wx"
		case '4':
			permissions+="r--"
		case '5':
			permissions+="r-x"
		case '6':
			permissions+="rw-"
		case '7':
			permissions+="rwx"
		}
	}
	n := Node{permissions, file.Nlink, file.Uid, file.Gid, file.Size,
	timeStr, file.Name}
	return n
}
