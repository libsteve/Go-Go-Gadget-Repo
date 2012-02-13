package ttt

import "strings"
import "./igame"

const(
	noPlayer = (games.Player)(iota)
	player1
	player2

)

type Game struct{
	Board [3][3]games.Player
	Choices map[string]int
}

func newGame() *Game{
	news := new(Game)	
	for _,rows := range news.Board {
		for j,_ := range rows {
			rows[j] = noPlayer
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

func(game *Game) MakeMove(player int, move string){
	n := game.Choices[strings.ToLower(move)]
	var p games.Player
	switch player{
		case 0:
			p = player1
		case 1:
			p = player2
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
	rps := strings.ToLower(move)
	if _, ok := game.Choices[rps]; ok { return true }
	return false
}

func (game *Game) Finished() (bool, games.Player){
	 for _, row := range game.Board{
	 	for _,square := range row{
	 		if square == noPlayer{
	 			return false, noPlayer
	 		}
	 	}
	 }
	 for _, row := range game.Board{
		if row[0] == row[1]  && row[1] == row[2]{
			return true, row[0]
		}
	 }
	 for i := 0; i < 3; i++{
	 	if game.Board[i][0] == game.Board[i][1] && game.Board[i][0] == game.Board[i][2]{
	 		return true, game.Board[i][0]
	 	}
	 }
	 if game.Board[0][0] == game.Board[1][1] &&game.Board[0][0] == game.Board[2][2]{
	 	return true, game.Board[0][0]
	 }
	 if game.Board[0][2] == game.Board[1][1] &&game.Board[0][2] == game.Board[2][0]{
	 	return true, game.Board[0][2]
	 }
	 return false, noPlayer
}

func (game *Game) getBoard() []string{
	var array []string
	for _,rows := range game.Board{
		str := ""
		for _,square := range rows{
			switch square{
				case noPlayer:
					str+="  "
				case player1:
					str+="x "
				case player2:
					str+="o "
			}
		}
		strings.TrimSpace(str)
		array = append(array,str)
	}
	return array
}