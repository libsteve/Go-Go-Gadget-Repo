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

type OutView struct {
	*View
	Reader *os.File
	Writer *os.File
}

/*
The run loop for the view.
*/
func (v *OutView) Loop() os.Error {

	var player_id string
	var opponent_move string
	for {
		
		req := <- v.Request

		switch req.Command {
		case Enable:
			// unfreeze user input/output
			// allow player to give input
			player_id = req.Args[0]
			v.enable(player_id)

		case Get:
			v.Response <- []string{v.get()} // the player's move
			// freeze user input/output

		case Set:
			// save the other player's move
			opponent_move = req.Args[0]

		case Show:
			// display the updated board
			player_id = req.Args[0]
			v.show(player_id, opponent_move)

		case Done:
			outcome := Outcome(req.Args[0])
			switch outcome {
			case Draw:
				// show that the game is a tie
				v.writeln(string(outcome))

			default:
				// show that the other player won
				v.writeln(player_id + " " + string(outcome) + "s")
			}
		}

	}

	return nil
}

func (v *OutView) write(msg string) {
	v.Writer.Write([]byte(msg))
}

func (v *OutView) writeln(msg string) {
	v.write(msg + "\n")
}

func (v *OutView) enable(player_id string) {
	v.write("Player " + player_id + "'s Move: ")
}

func (v *OutView) get() string {
	r := bufio.NewReader(v.Reader)
	raw, _, _ := r.ReadLine()
	return string(raw)
}

func (v *OutView) show(player_id string, opponent_move string) {
	v.writeln("Player " + player_id + "'s Opponent's Move: " + opponent_move)
}