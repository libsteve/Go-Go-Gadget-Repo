package main

import ( "http"; "flag"; "fmt"; "os" )

func main() {
	//////
	// create the flag arguments
	var port *string
	port = flag.String("port", ":8080", "The port to use for the server")
	flag.Parse()
	args := flag.Args()

	//////
	// make sure there are channels for the server
	if len(args) < 0 {
		fmt.Fprintln(os.Stderr, "Too Little Arguments")
		return
	}

	//////
	// create the map of channels for the server
	server := &server{make(map[string]chan string)}

	get := func(w http.ResponseWriter, r *http.Request) {
		server.getvalue(w, r)
	}

	set := func(w http.ResponseWriter, r *http.Request) {
		server.setvalue(w, r)
	}

	for _, named_channel := range args {
		server.channels[named_channel] = make(chan string), true

		getkey := "/?key="+named_channel
		http.HandleFunc(getkey, get)

		setkey := "/?key="+named_channel+"&value="
		http.HandleFunc(setkey, set)

		println(getkey)
		println(setkey)
	}

	http.ListenAndServe(*port, nil)
}

type server struct {
	channels map[string]chan string
}

func (s *server) setvalue(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if value, ok := values["key"]; ok {
		if channel, ok := s.channels[value[0]]; ok {
			if value, ok = values["value"]; ok {
				channel <- value[0]
				fmt.Println(w, "Success")
			}
		}
	}
}

func (s *server) getvalue(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if value, ok := values["key"]; ok {
		if channel, ok := s.channels[value[0]]; ok {
			fmt.Fprint(w, <-channel)
		}
	}
}