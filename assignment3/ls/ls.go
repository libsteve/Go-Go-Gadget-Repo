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
	"time"
	"container/vector"
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

type FileDatavector vector.Vector FileData
type DirDataVector vector.Vector FileDataVector






func Ls(dirname string, R bool, t bool) (DirDataVector, os.Err) {
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
type sortfunc func([]*os.FileInfo) []*os.FileInfo

// function to go through directories recursively
func recurdir(filename string, sort sortfunc) (DirDataVector, os.Err) {

	directories := new(DirDataVector)
	maindir := new(FileDataVector)
	dirqueue := new(FileDataVector)
	if fi, ok = os.Stat(filename); ok {
		if fi.IsDirectory() {

			maindir.Push(fileinfo(fi))

			files := ioutil.ReadDir(filename)
			files = sort(files)
			for index, file := range files {

				if file.IsDirectory() {
					morefiles := recurdir(file.Name, sort)
					if morefiles.Len > 0 {
						morefiles.Insert(0, file)
					}
					dirqueue.Push(morefiles)
				}

				maindir.Push(fileinfo(file))
			}
		}

		directories.Push(maindir)
		for index, dir := range *dirqueue {
			directories.Push(dir)
		}

		return directories, ok
	} else {
		return directories, ok
	}

}

// function not go recursively through directories
func readdir(filename string, sort sortfunc) (DirDataVector, os.Err) {

	directories := new(DirDataVector)
	maindir := new(FileDataVector)
	if fi, ok = os.Stat(filename); ok {
		if fi.IsDirectory() {

			maindir.Push(fileinfo(fi))

			files := ioutil.ReadDir(filename)
			files = sort(files)
			for index, file := range files {
				maindir.Push(fileinfo(file))
			}
		}

		directories.Push(maindir)

		return directories, ok
	} else {
		return directories, ok
	}

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
func alphasort(files []*os.FileInfo) []*os.FileInfo {
	sort.Sort(files.(alphaSort))
	files.([]*os.FileInfo)
}
// sort function to sort by timestamp
func timesort(files []*os.FileInfo) []*os.FileInfo {
	sort.Sort(files.(timeSort))
	files.([]*os.FileInfo)
}






// display function to display all information
func fileinfo(file *os.FileInfo) FileDataNode {
	node := fileInfoToNode(file)
	return node
}

// converts the *os.FileInfo to FileData
func fileInfoToNode(file *os.FileInfo) FileData{
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
	n := FileData{permissions, file.Nlink, file.Uid, file.Gid, file.Size,
	timeStr, file.Name}
	return n
}
