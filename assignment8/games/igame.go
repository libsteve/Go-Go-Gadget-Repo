package games

type Player int

const(
	NO_PLAYER = (Player)(iota)
	PLAYER_1
	PLAYER_2

)

/*
An interface for a game.
*/
type Igame interface {
	

	IsSimultaneous() bool
	/*
	Check that the given move is a valid move.

	Parameters:
		move - a move string
	
	Returns:
		true if the move is valid, false otherwise
	*/
	CheckMoveValid(move string) bool
	
	/*
	Make the given move for the given player.

	Parameters:
		player - the player id (id starts from 0)
		move - a valid move string
	*/
	MakeMove(player int, move string)

	/*
	Check to see if the game is finished.

	Returns:
		true if the game is finished, false otherwise
		the player that won or noplayer if there is a tie or the game isn't finished yet
	
	*/
	Finished() (bool, Player)

}