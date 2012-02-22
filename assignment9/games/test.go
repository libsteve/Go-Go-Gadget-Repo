package main

import ( "http"; "flag"; "fmt"; "os"; "bufio"; "strings"; "url" )

func main() {
	var host *string
	host = flag.String("host", "localhost:8080", "The server host name and port to connect to (defaults to localhost:8080)")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 { fmt.Fprintln(os.Stderr, "Not Enough Args."); return }

	client := http.DefaultClient

	for _, arg := range args {
		//////
		// get the response for the query
		var vals url.Values
		if vals = parse(arg); vals == nil { continue }
	//	if response, err := client.Get("http://" + *host + query); err == nil {
		if response, err := client.PostForm("http://" + *host, vals); err == nil{
			//////
			// read and print the response
			r := bufio.NewReader(response.Body)
			ln, _, _ := r.ReadLine()
			fmt.Println(string(ln))

			//////
			// close the response readcloser
			response.Body.Close()

		}
	}
}

/*
parse the argument and return the proper query format
*/
func parse(arg string) url.Values {
	vals := make(url.Values)
	split := strings.Split(arg, "=")
	if len(split) < 1 || len(split) > 2 {
		fmt.Fprintln(os.Stderr, arg + " : invalid argument; skipping")
		return nil
	}else{
		if len(split) == 1{
			vals.Add("key", split[0])
		}else{
			vals.Add("key", split[0])
			vals.Add("value", split[1])
		}
	}
	return vals
}