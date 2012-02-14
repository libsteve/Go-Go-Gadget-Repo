package games

import ( "os"; "bufio" )

/*
Create a new View struct.
*/
func NewView() *View {
	v := new(View)
	v.Request = make(chan Request)
	v.Response = make(chan Response)
	return v
}

/*
The run loop for the view.
*/
func (v *View) Loop() os.Error {

	var player_id string
	var opponent_move string
	for {
		
		req := <- v.Request

		switch req.Command {
		case Enable:
			// unfreeze user input/output
			// allow player to give input
			player_id = req.Args[0]
			enable(player_id)

		case Get:
			v.Response <- []string{get()} // the player's move
			// freeze user input/output

		case Set:
			// save the other player's move
			opponent_move = req.Args[0]

		case Show:
			// display the updated board
			player_id = req.Args[0]
			show(player_id, opponent_move)

		case Done:
			outcome := Outcome(req.Args[0])
			switch outcome {
			case Draw:
				// show that the game is a tie
				println(string(outcome))

			default:
				// show that the other player won
				println(player_id + " " + string(outcome) + "s")
			}
		}

	}

	return nil
}

func enable(player_id string) {
	print("Player " + player_id + "'s Move: ")
}

func get() string {
	r := bufio.NewReader(os.Stdin)
	raw, _, _ := r.ReadLine()
	return string(raw)
}

func show(player_id string, opponent_move string) {
	println("Player " + player_id + "'s Opponent's Move: " + opponent_move)
}