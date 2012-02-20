package ttt

import ( "strings"; "./games" )

/*
 The Game struct

 Has a board and choices 
 */
type Game struct{
	Board [3][3]games.Player
	Choices map[string]int
}

/*
 Creates a new game
 */
func NewGame() *Game{
	news := new(Game)	
	for i,_ := range news.Board {
		for j,_ := range news.Board[i] {
			news.Board[i][j] = games.NO_PLAYER
		}
	}
	news.Choices = map[string]int {
		"nw" : 0,
		"n"  : 1,
		"ne" : 2,
		"w"  : 3,
		"c"  : 4,
		"e"  : 5,
		"sw" : 6,
		"s"  : 7,
		"se" : 8 }
	return news
}

/*
 Makes a move
 */
func(game *Game) MakeMove(player int, move string){
	n := game.Choices[strings.ToLower(move)]
	var p games.Player
	switch player{
		case 0:
			p = games.PLAYER_1
		case 1:
			p = games.PLAYER_2
	}
	game.Board[n/3][n%3] = p
}

/*
Take a string and check to see if it is a valid RPS move.

Parameters:
	move - a string that is the pleyer's move

Returns:
	bool - true if the move is valid, false otherwise
*/
func (game *Game) CheckMoveValid(move string) bool {
	ttt := strings.ToLower(move)
	if _, ok := game.Choices[ttt]; ok { 
		n := game.Choices[ttt]
		if(game.Board[n/3][n%3] != games.NO_PLAYER){
			return false
		}
		return true 
	}
	return false
}

/*
 Is the game finished

 returns a bool representing whether the game is finished and the player who one (no player if draw)
 */
func (game *Game) Finished() (bool, games.Player){
	 for _, row := range game.Board{
		if row[0] == row[1]  && row[1] == row[2] && row[0] != games.NO_PLAYER{
			return true, row[0]
		}
	 }
	 for i := 0; i < 3; i++{
	 	if game.Board[0][i] == game.Board[1][i] && game.Board[0][i] == game.Board[2][i] && game.Board[0][i] != games.NO_PLAYER{
	 		return true, game.Board[i][0]
	 	}
	 }
	 if game.Board[0][0] == game.Board[1][1] &&game.Board[0][0] == game.Board[2][2] && game.Board[0][0] != games.NO_PLAYER{
	 	return true, game.Board[0][0]
	 }
	 if game.Board[0][2] == game.Board[1][1] &&game.Board[0][2] == game.Board[2][0] && game.Board[0][2] != games.NO_PLAYER{
	 	return true, game.Board[0][2]
	 } 
	 for _, row := range game.Board{
	 	for _,square := range row{
	 		if square == games.NO_PLAYER{
	 			return false, games.NO_PLAYER
	 		}
	 	}
	 }
	 return true, games.NO_PLAYER
}

/*
 Are player moves simultaneous
 */
func (game *Game) IsSimultaneous() bool {
	return false;
}

/*
 Clears the board
 */
func (game *Game) Clear(){
	for i,_ := range game.Board {
		for j,_ := range game.Board[i] {
			game.Board[i][j] = games.NO_PLAYER
		}
	}
}