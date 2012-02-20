package main

import ( "./games"; "./rps"; "./ttt" )
import ( "flag"; "fmt"; "os"; "netchan" )

func main() {
	var r *bool
	var t *bool
	r = flag.Bool("rps", false, "Play a game of Rock-Paper-Scissors")
	t = flag.Bool("ttt", false, "Play a game of Tick-Tack-Toe")
	flag.Parse()

	args := flag.Args()
	var port string
	if len(args) == 1 {
		port = args[0]
	} else {
		fmt.Fprint(os.Stderr, "You must provide a port to connect to.")
		return
	}

	if *r && *t {
		fmt.Fprintln(os.Stderr, "You can only play one game at a time.")
		return
	} else if !( *r || *t ) {
		// connext to an exporter with an importer
		// run the importer loop
		v2 := games.NewView()

		importer, err := netchan.Import("tcp", "localhost" + port)
		if err != nil { fmt.Fprintln(os.Stderr, err.String()); return }

		err = importer.Import("Request", v2.Request, netchan.Recv, 1)
		if err != nil { fmt.Fprintln(os.Stderr, err.String()); return }

		err = importer.Import("Responce", v2.Response, netchan.Send, 1)
		if err != nil { fmt.Fprintln(os.Stderr, err.String()); return }

		ov2 := &games.ViewView{v2, os.Stdin, os.Stdout}

		ov2.Loop()
		
		return
	}

	var game games.Igame

	if *r {
		game = rps.NewGame()
	} else if *t {
		game = ttt.NewGame()
	}

	v1 := games.NewView()
	v2 := games.NewView() // <- this will be the exported/inported thingy
	ref := games.NewReferee(game, v1, v2)

	exporter := netchan.NewExporter()
	exporter.Export("Request", v2.Request, netchan.Send)
	exporter.Export("Responce", v2.Response, netchan.Recv)

	exporter.ListenAndServe("tcp", port)

	ov1 := &games.ViewView{v1, os.Stdin, os.Stdout}

	go ov1.Loop()

	ref.Loop()
	return
}
