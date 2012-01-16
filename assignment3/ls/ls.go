/*
A group of functions related to listing all files in a given directory.

ls.Ls() takes 2 parameters:
	1) the directory name
	2) a list of parameters

Parameters can be any of the following:
	none - list one line at a time, display files in alphabetical order, no recursive directory searching
	"-n" - display with information
	"-R" - go through directories recursively
	"-t" - sort by timestamp

Usage:
	ls.Ls(directoryname, argumentlist)
 */
package ls

import (
	"io/ioutil"
	"os"
	"time"
	"sort"
	"fmt"
)





/*
A struct that represents the data of a file

Values:
	Mode string	-	permissions
	Nlink uint64	-	hardlink count
	Uid int		-	user id
	Gid int		-	group id
	Size int64	-	byte size
	Mtime string	-	last modified time
	Name string	-	file/folder name
	Blocks int64	-	number of blocks
 */
type FileData struct{
	Mode string
	Nlink uint64
	Uid int
	Gid int
	Size int64
	Mtime string
	Name string
	Blocks int64
}





/*
Find the files and folders in the given path.

Parameters:
	dirname string	-	the path of the first file/folder
	R bool		-	whether or not the search will be recursive
	t bool		-	whether or not files are sorted by name or by time

Returns:
	[][]FileData	-	an array of arrays of ls.FileData structs
						[n][0] is the directory that the [n] array represents
						[n][x] is the file/folder in directory n

Usage:
	fileData := ls.Ls("../..", true, true)
		search thorugh path "../.." with recursion and sort by time

 */
func Ls(dirname string, R bool, t bool) ([][]FileData, os.Error) {
	lsdir := readdir
	sort := alphasort

	if R {
		lsdir = recurdir
	}
	if t {
		sort = timesort
	}

	return lsdir(dirname, sort)
}






// the sort function type
type sortfunc func([]*os.FileInfo)

// function to go through directories recursively
func recurdir(filename string, sort sortfunc) ([][]FileData, os.Error) {
	var ok os.Error
	directories := make([][]FileData, 0)
	maindir := make([]FileData, 0)
	dirqueue := make([][]FileData, 0)
	if fi, ok := os.Stat(filename); ok == nil {
		if fi.IsDirectory() {

			maindir = append(maindir, fileinfo(fi))

			if files, ok2 := ioutil.ReadDir(filename); ok2 == nil {
				sort(files)
				for _, file := range files {

					if file.IsDirectory() {
						if moredirs, ok3 := recurdir(filename + "/" + file.Name, sort); ok3 == nil {
							for _, morefiles := range moredirs {
								dirqueue = append(dirqueue, morefiles)
							}
						} else {
							return directories, ok3
						}
					}

					maindir = append(maindir, fileinfo(file))
				}
			} else {
				return directories, ok2
			}
		} else {
			root, _ := os.Stat("./")
			maindir = append(maindir, fileinfo(root))
			maindir = append(maindir, fileinfo(fi))
		}

		directories = append(directories, maindir)
		for _, dir := range dirqueue {
			directories = append(directories, dir)
		}
		return directories, ok
	}
	return directories, ok

}

// function not go recursively through directories
func readdir(filename string, sort sortfunc) ([][]FileData, os.Error) {
	var ok os.Error
	directories := make([][]FileData, 0)
	maindir := make([]FileData, 0)
	if fi, ok := os.Stat(filename); ok == nil {
		if fi.IsDirectory() {

			maindir = append(maindir, fileinfo(fi))

			if files, ok2 := ioutil.ReadDir(filename); ok2 == nil {
				sort(files)
				for _, file := range files {
					maindir = append(maindir, fileinfo(file))
				}

			} else {
				return directories, ok2
			}
		} else {
			root, _ := os.Stat("./")
			maindir = append(maindir, fileinfo(root))
			maindir = append(maindir, fileinfo(fi))
		}

		directories = append(directories, maindir)

		return directories, ok
	}

	return directories, ok

}





type alphaSort []*os.FileInfo
func (s alphaSort) Len() int {
	return len(s)
}
func (s alphaSort) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
func (s alphaSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type timeSort []*os.FileInfo
func (s timeSort) Len() int {
	return len(s)
}
func (s timeSort) Less(i, j int) bool {
	return s[i].Mtime_ns < s[j].Mtime_ns
}
func (s timeSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort function to sort alphabetically
func alphasort(files []*os.FileInfo) {
	sort.Sort((alphaSort)(files))
}
// sort function to sort by timestamp
func timesort(files []*os.FileInfo) {
	sort.Sort((timeSort)(files))
}






// display function to display all information
func fileinfo(file *os.FileInfo) FileData {
	node := fileInfoToNode(file)
	return node
}

// converts the *os.FileInfo to FileData
func fileInfoToNode(file *os.FileInfo) FileData{
	t := time.NanosecondsToLocalTime(file.Mtime_ns);
	timeStr := t.Format("Jan _2 15:04");
	permissions :=""
	permo := fmt.Sprintf("%o", file.Mode)
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
	n := FileData{permissions, file.Nlink, file.Uid, file.Gid, file.Size,
	timeStr, file.Name, file.Blocks}
	return n
}
