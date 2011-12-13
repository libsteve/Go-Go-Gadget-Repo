package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"bufio"
	"http"
//	"./prepro"
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

func handlerInput(w http.ResponseWriter, r *http.Request){
	//input := r.FormValue("input")
	/*take the input and stick it in your thing*/
}

func main() {
	var ws *bool
	ws = flag.Bool("ws",false, "run as a webserver")
	flag.Parse()
	if *ws {
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
		file:= os.Open("temp.txt")
		in := bufio.NewReader(file)
		/* read from file in */
		file.Close()
		os.Remove("temp.txt")
		for _, arg := range flag.Args() {
			if arg != "-"{
				file,_ := os.Open(arg)
				in := bufio.NewReader(file)
				/* read from this file*/
				file.Close()
			}
		}
	}
}
