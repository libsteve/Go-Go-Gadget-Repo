package games

import ( "os"; "bufio" )

/*
Create a new View struct.
*/
func NewView() *games.View {
	v := new(View)
	v.Request = make(chan games.Request)
	v.Response = make(chan games.Response)
	return v
}

/*
The run loop for the view.
*/
func (v *games.View) Loop() os.Error {

	var req Request

	var player_id string
	var opponent_move string
	for {
		
		req <- v.Request

		switch req.Command {
		case Enable:
			// unfreeze user input/output
			// allow player to give input
			player_id = req.Args[0]
			enable(player_id)

		case Get:
			v.Responce <- get() // the player's move
			// freeze user input/output

		case Set:
			// save the other player's move
			opponent_move = req.Args[0]

		case Show:
			// display the updated board
			show(player_id, opponent_move)

		case Done:
			outcome := games.Outcome(req.Args[0])
			switch outcome {
			case Draw:
				// show that the game is a tie
				println(string(outcome))

			default:
				// show that the other player won
				println(player_id " " + string(outcome) + "s")
			}
		}

	}
}

func enable(player_id string) {
	print(player_id + "'s' Move: ")
}

func get() string {
	r := bufio.NewReader(os.Stdin)
	raw, _, _ := Reader.ReadLine()
	return string(raw)
}

func show(player_id string, opponent_move string) {
	println(player_id + "'s' Opponent's Move: " + opponent_move)
}