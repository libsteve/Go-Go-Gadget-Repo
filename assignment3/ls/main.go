package main

import (
	"os"
	"strings"
	"flag"
	"fmt"
	"template"
	"./ls"
)
/*
 * Run the program
 *
 * Usage: 
 *		./ls [args] directory_name
 *		
 * 		possible args:
 *		-R: go through directories recursively
 *		-n: print with information
 *		-t: sort files by modification time
 *
 *		if no arguments are getting, print out alphabetically with 1 file
 *		per line 
 */
func main() {
	var R *bool
	var n *bool
	var t *bool

	R = flag.Bool("R", false, "go through directories recursively")
	n = flag.Bool("n", false, "print with information")
	t = flag.Bool("t", false, "sort files by modification time")
	flag.Parse()

	temp := template.Must(template.New("ls").Parse("{{.Mode}} {{printf `%3d` .Nlink}} {{.Uid}}  {{.Gid}}  {{printf `%7d` .Size}} {{.Mtime}}  {{.Name}}\n"))
	data, _ := ls.Ls(flag.Arg(0), *R, *t)
	path := flag.Arg(0)
	if strings.HasSuffix(path, "/"){
		path = path[0: len(path)-1]
	}
	printFiles(data, path, n, temp)
}

/*
 * Calculate the total blocks in a directory
 */
func totalBlocks(dir []ls.FileData) int64{
	var total int64
	for pos, file := range dir{
		if(pos != 0){
			total+= file.Blocks
		}
	}
	return total
}

/*
 * Prints out the information given to you by ls.Ls depending on the flags
 * given
 *
 */
func printFiles(data [][]ls.FileData, path string, n *bool, temp *template.Template) {
	for pos, dir := range data {
		if pos != 0{
			path+="/"
			path+= dir[0].Name
			fmt.Printf("\n%s:\n", path)
		}
		if(*n){
			fmt.Printf("total %d\n", totalBlocks(dir))
		}
		for pos1, file := range dir{
			if pos1 != 0{
				if (*n){
					temp.Execute(os.Stdout, file)
				} else{
					fmt.Println(file.Name)
				}
			}
		}
	}

}
