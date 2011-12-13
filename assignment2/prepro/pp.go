package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"http"
	"./prepro"
)
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
func main() {
	var ws *bool
	ws = flag.Bool("ws",false, "run as a webserver")
	flag.Parse()
	if *ws {
		http.HandleFunc("/", handler)
   		http.ListenAndServe(":6060", nil)
		r, err := http.Get("http://localhost:6060/pp")
	}else if flag.NArg() == 0{
		in := bufio.NewReader(os.Stdin)
		/* read from standard in */

	} else {
		for _, arg := range flag.Args() {
			if arg == "-" {
				in := bufio.NewReader(os.Stdin)
				/* read from standard in */
			} else{
				file,_ := os.Open(arg)
				in := bufio.NewReader(file)
				/* read from this file*/
				file.Close()
			}
		}	
	}
}
