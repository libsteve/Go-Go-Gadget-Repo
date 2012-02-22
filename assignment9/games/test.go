package main

import ( "http"; "net"; "flag"; "fmt"; "os" )

func main() {
	var host *string
	host = flag.String("host", "localhost:8080", "The server host name and port to connect to (defaults to localhost:8080)")
	flag.Parse()

	if conn, err := net.Dail("tcp", *host); err != nil {
		fmt.Fprintln(os.Stderr, err.String())
		return
	}

	http.NewClientConn()
	
}