package ttt

import "strings"

const (
	choices = map[string]int {
		"nw" : 0
		"n"  : 1
		"ne" : 2
		"w"  : 3
		"c"  : 4
		"e"  : 5
		"sw" : 6
		"s"  : 7
		"se" : 8
	}

	noPlayer = (Player)(iota)
	player1
	player2

)

type Player int

type Game struct{
	Board int[3][3]
}

func newGame() *Game{
	news := new(Game)	
	for _, row := range news.Board {
		for _ , square := range row {
			square = 0
		}
	}
	return news
}

func(game *Game) MakeMove(player int, move string){
	n := choices[strings.ToLower(move)]
	switch player
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
	if _, ok := choices[rps]; ok { return true }
	return false
}

func  