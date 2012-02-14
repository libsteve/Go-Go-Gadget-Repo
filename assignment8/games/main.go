package main

import ( "./games"; "./rps"; "./ttt"; "flag"; "fmt"; "os" )

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

	v1 := games.NewView()
	v2 := games.NewView()
	ref := games.NewReferee(game, v1, v2)

	go v1.Loop()
	go v2.Loop()

	ref.Loop()

}