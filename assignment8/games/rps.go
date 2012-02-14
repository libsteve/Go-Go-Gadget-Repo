package rps

import ( "strings"; "./games" )

/*
 Creates a new rock paper scissor game
 */
type Game struct{
	player1Move int
	player2Move int
	Choices map[string]int
}

/*
 Create a new game
 */
func NewGame() *Game{
	news := new(Game)
	news.Choices = map[string]int{
		"rock"		:	1 ,
		"paper"		:	2 ,
		"scissors"	:	3 }
	news.player1Move = 0
	news.player2Move = 0
	return news
}

/*
 Make a move 
 */
func(game *Game) MakeMove(player int, move string){
	n := game.Choices[strings.ToLower(move)]
	if player == 0{
		game.player1Move = n
	} else if player == 1 {
		game.player2Move = n
	}
}

/*
Take a string and check to see if it is a valid RPS move.

Parameters:
	move - a string that is the pleyer's move

Returns:
	bool - true if the move is valid, false otherwise
*/
func (game *Game) CheckMoveValid(move string) bool {
	rps := strings.ToLower(move)
	if _, ok := game.Choices[rps]; ok { return true }
	return false
}

/*
 Is the game finished

 returns a bool representing whether the game is finished and the player who one (no player if draw)
 */
func (game *Game) Finished() (bool, games.Player){
	if game.player1Move == 0 || game.player2Move == 0 {
		return false, games.NO_PLAYER
	} else if game.player1Move - game.player2Move == 0{
		return true, games.NO_PLAYER
	} else if ((game.player1Move - game.player2Move)%3+3)%3 == 1 { 
		return true, games.PLAYER_1
	}
	return true, games.PLAYER_2
}

/*
 Are player moves simultaneous
 */
func (game *Game) IsSimultaneous() bool {
	return true;
}

/*
 Clears the board
 */
func (game *Game) Clear(){
	game.player1Move = 0
	game.player2Move = 0
}
