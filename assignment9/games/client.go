package main

import ( "./games"; "./rps"; "./ttt" )
import ( "http"; "url"; "bufio"; "flag"; "fmt"; "os" )

func main() {
	var r *int
	var t *int
	r = flag.Int("rps", 0, "Play a game of Rock-Paper-Scissors as Player 1 or 2")
	t = flag.Int("ttt", 0, "Play a game of Tick-Tack-Toe as Player 1 or 2")

	var host *string
	host = flag.String("host", "localhost:8080", "The server host name and port to connect to")
	flag.Parse()


	var player_id int
	var game games.Igame
	if *r <= 0 && *t <= 0 {
		fmt.Fprintln(os.Stderr, "A game and a player number must be provided")
		return
	} else if *r > 2 || *t > 2 {
		fmt.Fprintln(os.Stderr, "There can only be a Player 1 or a Player 2")
		return
	} else if (*r > 0 && *r < 3) && (*t > 0 && *t < 3) {
		fmt.Fprintln(os.Stderr, "Only one game can be played at a time")
		return
	}
	if *r != 0 {
		player_id = *r
		game = rps.NewGame()
	} else { 
		player_id = *t
		game = ttt.NewGame()
	}


	// play the games
	var ref *games.Referee
	player := games.NewView()
	player_view := &games.ViewView{player, os.Stdin, os.Stdout}
	var prox *proxy
	prox_back := games.NewView()
	prox = &proxy{prox_back, *host, "", ""}
	prox.url = *host

	switch player_id {
	case 1:
		prox.player_id = "q1"
		prox.proxy_id = "q2"
		ref = games.NewReferee(game, player, prox_back)
	case 2:
		prox.player_id = "q2"
		prox.proxy_id = "q1"
		ref = games.NewReferee(game, prox_back, player)
	}

	go player_view.Loop()
	go loop(prox)

	ref.Loop()
}

type proxy struct {
	*games.View
	url string
	proxy_id string
	player_id string
}

func loop(prox *proxy) {
	client := http.DefaultClient
	for {
		request := <- prox.Request

		switch request.Command {
		case games.Get:
			go func(request games.Request) {
				vals := make(url.Values)
				vals.Add("key", prox.player_id)
				if response, err := client.PostForm("http://" + prox.url, vals); err == nil{
					//////
					// read and print the response
					r := bufio.NewReader(response.Body)
					ln, _, _ := r.ReadLine()
					resp := string(ln)

					//////
					// close the response readcloser
					response.Body.Close()
					prox.Response <- []string{ resp }
				}
			}(request)
		case games.Set:
			go func(request games.Request) {
				vals := make(url.Values)
				vals.Add("key", prox.proxy_id)
				vals.Add("value", request.Args[0])
				client.PostForm("http://" + prox.url, vals)
			}(request)
		default:
			continue
		}
	}
}