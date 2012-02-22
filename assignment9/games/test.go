package main

import ( "http"; "flag"; "fmt"; "os";  "strings" )

func main() {
	var host *string
	host = flag.String("host", "localhost:8080", "The server host name and port to connect to (defaults to localhost:8080)")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 { fmt.Fprintln(os.Stderr, "Not Enougn Args."); return }

	client := http.DefaultClient

	for _, arg := range args {
		//////
		// get the response for the query
		var query string
		if query = parse_arg(arg); query == "" { continue }
		if response, err := client.Get("http://" + *host + query); err != nil {

			//////
			// read and print the response
			//r := bufio.NewReader(response.Body.(io.Reader))
			//ln, _, _ := r.ReadLine()
			//fmt.Println(string(ln))
			this := []byte("                                                     ")
			response.Body.Read(this)
			fmt.Println(string(this))

			//////
			// close the response readcloser
			response.Body.Close()

		}
	}
}

/*
parse the argument and return the proper query format
*/
func parse_arg(arg string) string {
	key := "/?key="
	value := "&value="

	split := strings.Split(arg, "=")
	if len(split) < 1 || len(split) > 2 {
		fmt.Fprintln(os.Stderr, arg + " : invalid argument; skipping")
	} else {
		if len(split) == 1 {
			return key + split[0]
		} else {
			return key + split[0] + value + split[1]
		}
	}
	return ""
}