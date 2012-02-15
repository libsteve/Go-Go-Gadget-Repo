package games

import ( "strconv"; "os" )

/*
A struct to represent a referee handling many players' views for a given game.
*/
type Referee struct {
	players []*View
	game Igame
}

/*
Create a new Referee struct.

Parameters:
	game - the game
	views - all the views for every player
*/
func NewReferee(game Igame, views ...*View) *Referee {
	r := new(Referee)
	r.players = views
	r.game = game
	return r
}

/*
Loop throught the Referee's operations.
*/
func (r *Referee) Loop() os.Error {
	var simultaneous bool
	simultaneous = r.game.IsSimultaneous()

	for {

		var move string
		
		for id, player := range r.players {

			///////
			// check to see if the game is finished
			if r.checkFinished() {
				defer r.Loop()
				return nil
			}

			///////
			// repeat this until the user's move is valid
			for {
				///////
				// enable the player's view
				player.Request <- Request{Enable, []string{strconv.Itoa(id+1)}}

				///////
				// get the player's input
				player.Request <- Request{Get, []string{}}

				///////
				// check to see if the player's move is valid
				response := <- player.Response
				if move = string(response[0]); r.game.CheckMoveValid(move) {
					r.game.MakeMove(id,move)
					break
				}
			}

			///////
			// set this player's move for all other players
			for _, other := range r.players {
				if other != player { 
					other.Request <- Request{Set, []string{move}} 
				}
			}
			
			if !simultaneous { r.show(id) }			

		}

		if simultaneous { r.show() }

	}

	return nil
}

/*
Check to see if the game is finished.

If the game is finished, notify all players and then return true, otherwise return false.
*/
func (r *Referee) checkFinished() bool {
	if finished, winner := r.game.Finished(); finished {
		for id, player := range r.players {
			if winner == NO_PLAYER {
				player.Request <- Request{Done, []string{string(Draw)}}
			} else if winner == Player(id+1) {
				player.Request <- Request{Done, []string{string(Win)}}
			} else {
				player.Request <- Request{Done, []string{string(Lose)}}
			}
		}
		r.game.Clear()
		return true
	}
	return false
}

/*
Have all players show the move.
*/
func (r *Referee) show(except_ids ...int) {
	except := make(map[int]bool)
	for _, other_id := range except_ids { 
		except[other_id] = true, true
	}
	for id, player := range r.players {
		if _, ok := except[id]; ok { continue }
		player.Request <- Request{Show, []string{strconv.Itoa(id+1)}}
	}
}