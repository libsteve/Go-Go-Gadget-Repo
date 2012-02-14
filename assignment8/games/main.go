package main

import ( "./games"; "./rps"; "./ttt" )
import ( "flag"; "fmt"; "os" )

func main() {

	var r *bool
	var t *bool
	r = flag.Bool("rps", false, "Play a game of Rock-Paper-Scissors")
	t = flag.Bool("ttt", false, "Play a game of Tick-Tack-Toe")
	flag.Parse()

	if *r && *t { 
		fmt.Fprintln(os.Stderr, "You can only play one game at a time.")
		return
	} else if !( *r || *t ) {
		fmt.Fprintln(os.Stderr, "You must select a game to play.")
		return
	}

	var game games.Igame

	if *r {
		game = rps.NewGame()
	} else if *t {
		game = ttt.NewGame()
	}

	args := flag.Args()
	var file *os.File
	if len(args) == 1 { 
		var err os.Error
		if file, err = os.OpenFile(args[0], os.O_RDWR, 0666); err != nil {
			fmt.Fprintln(os.Stderr, "404: File Not Found.")
			return
		}
	}

	v1 := games.NewView()
	v2 := games.NewView()
	ref := games.NewReferee(game, v1, v2)

	ov1 := &games.ViewView{v1, os.Stdin, os.Stdout}

	var ov2 *games.ViewView
	if file != nil {
		ov2 = &games.ViewView{v2, file, file}
	} else {
		ov2 = &games.ViewView{v2, os.Stdin, os.Stdout}
	}

	go ov1.Loop()
	go ov2.Loop()

	ref.Loop()

}