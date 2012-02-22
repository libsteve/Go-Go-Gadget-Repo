package main

import ( "./games"; "./rps"; "./ttt" )
import ( "http"; "flag"; "fmt"; "os" )

func main() {
	var rps *string
	var ttt *string
	rps = flag.String("rps", 0, "Play a game of Rock-Paper-Scissors as Player 1 or 2")
	ttt = flag.String("ttt", 0, "Play a game of Tick-Tack-Toe as Player 1 or 2")

	var host *string
	host = flag.String("host", "localhost:8080", "The server host name and port to connect to")
	flag.Parse()


	var player_id int
	var game games.Igame
	if *rps <= 0 && *ttt <= 0 {
		fmt.Fprintln(os.Stderr, "A game and a player number must be provided")
		return
	} else if *rps > 2 || *ttt > 2 {
		fmt.Fprintln(os.Stderr, "There can only be a Player 1 or a Player 2")
		return
	} else if (*rps > 0 && *rps < 3) && (*ttt > 0 && *ttt < 3) {
		fmt.Fprintln(os.Stderr, "Only one game can be played at a time")
		return
	} else {
		if *rps != 0 {
			player_id = *rps
			game = rps.NewGame()
		} else { 
			player_id = *ttt
			game = ttt.NewGame()
		}
	}


	// play the games
	var ref games.Referee
	player := games.NewView()
	proxy := (&proxy)(games.NewView())
	proxy.url = *host

	switch player_id {
	case 1:
		ref := games.NewReferee(game, player, proxy)
	case 2:
		ref := games.NewReferee(game, proxy, player)
	}

	go player.Loop()
	go loop(proxy)

	ref.Loop()
}

type proxy struct {
	games.View
	url string
}

func loop(proxy *proxy) os.Error {
	for {
		request <- proxy.Request

		switch request.Command {
		case games.GET:
			//send get request
		case games.SEND:
			//sned send request
		default:
			continue
		}
	}
}