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

	handler := func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")
		if value != "" { server.setvalue(w, key, value) }
		if value == "" { server.getvalue(w, key) }
	}

	for _, named_channel := range args {
		server.channels[named_channel] = make(chan string), true

		http.HandleFunc("/", handler)
	}

	http.ListenAndServe(*port, nil)
}

type server struct {
	channels map[string]chan string
}

func (s *server) setvalue(w http.ResponseWriter, key, value string) {
	if channel, ok := s.channels[key]; ok {
		channel <- value
	}
}

func (s *server) getvalue(w http.ResponseWriter, key string) {
	if channel, ok := s.channels[key]; ok {
		fmt.Fprint(w, <-channel)
	}
}