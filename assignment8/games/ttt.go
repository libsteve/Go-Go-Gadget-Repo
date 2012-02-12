package ttt

import "strings"

const (
	choices = map[string]int {
		"n"		:	0
		"ne"	:	1
		"e"		:	2
		"se"	:	3
		"s"		:	4
		"sw"	:	5
		"w"		:	6
		"nw"	:	7
	}
)

/*
Take a string and check to see if it is a valid RPS move.

Parameters:
	move - a string that is the pleyer's move

Returns:
	bool - true if the move is valid, false otherwise
*/
func CheckMoveValid(move string) bool {
	rps := strings.ToLower(move)
	if _, ok := choices[rps]; ok { return true }
	return false
}