package games

/*
An interface for a game.
*/
type Igame interface {
	
	/*
	Check that the given move is a valid move.

	Parameters:
		move - a move string
	
	Returns:
		bool - true if the move is valid, false otherwise
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
		bool - true if the game is finished, flase otherwise
	*/
	Finished() bool

	/*
	Get the winner of the game.

	Returns:
		int - the player id of the winner, -1 if there is no winner
	*/
	Winner() int

}