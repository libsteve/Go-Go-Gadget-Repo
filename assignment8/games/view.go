package view

import ( "games" )

func (v *games.View) Loop() os.Error {
	var req Request
	for {
		
		req <- v.Request

		switch req.Command {
		case games.Enable:
			// allow player to give input
		case games.Get:
			v.Responce <- // the player's move
		case games.Set:
			// save the other player's move
		case games.Show:
			// display the updated board
		case games.Done:
			switch req.Args[0].(games.Outcome) {
			case games.Draw:
				// show that the game is a tie
			case games.Win:
				// show that this player won
			case games.Lose:
				// show that the other player won
			}
		}

	}
}