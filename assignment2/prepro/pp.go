/*
 * A simple preprocessor that takes in input files and can
 * 
 * 	include files,
 * 	(un-)define names,
 *	 and conditionally ex- or include text if a name is (un-)defined.
 * 
 * Command line Usage:
 *		./pp [-flags] [args]
 */
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"bufio"
	"http"
	"strings"
	"./prepro"
)

/*
 * The handler for the web server form
 * 
 * Parameters:
 *		w  -  a http response writer
 *		r  -  a http request header
 */
func handler(w http.ResponseWriter, r *http.Request) {
      file, _:= os.Open("index.html")
	  in:= bufio.NewReader(file)
	  for{
		line,err := in.ReadString('\n')
		if err != nil {
			break
		} else {
			fmt.Fprint(w, line)
		}
	  }
}
/*
 * The handler for the output after you preprocess the input
 * 
 * Parameters:
 *		w  -  a http response writer
 *		r  -  a http request header
 */
func handlerInput(w http.ResponseWriter, r *http.Request){
	input := r.FormValue("input")
	if strings.Contains(input, "#include") || strings.Contains(input, "# include") {
		fmt.Fprint(w, "Cannot include files on the web server")
	}else{
		ioutil.WriteFile("temp.txt", []uint8(input), 0600)
		file, _:= os.Open("temp.txt")
		in := bufio.NewReader(file)
		prepro.ReadInput(in)
		file.Close()
		os.Remove("temp.txt")
	}
}

/*
 * Run the program
 *
 * Usage:
 * 		./pp -       					: read from standard in
 *		./pp - file1.txt file2.txt ...  : read from standard in as well as files
 *		./pp -h							: launch the webserver
 *		./pp file1.txt file2.txt ...	: read from files
 *		./pp 							: read from standard in 
 */
func main() {
	var h *bool
	h = flag.Bool("h",false, "run as a webserver")
	flag.Parse()
	if *h {
		http.HandleFunc("/", handler)
		http.HandleFunc("/pp", handlerInput)
		http.ListenAndServe(":6060", nil)
	} else if flag.Arg(0) == "-" || flag.NArg() == 0{
		sin := bufio.NewReader(os.Stdin)
		b :=  []byte{}
		for {
			l,_,_ := sin.ReadLine()
			if string(l) == "."{
				ioutil.WriteFile("temp.txt",b , 0600)
				break
			} else{
				for _, a := range l{
					b = append(b,a)
				}
				b = append(b,'\n')
			}
		}
		file, _:= os.Open("temp.txt")
		in := bufio.NewReader(file)
		prepro.ReadInput(in)
		file.Close()
		os.Remove("temp.txt")
		for _, arg := range flag.Args() {
			if arg != "-"{
				file,_ := os.Open(arg)
				in := bufio.NewReader(file)
				prepro.ReadInput(in)
				file.Close()
			}
		}
	}
}
