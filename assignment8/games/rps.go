package rps

import "strings"

const (
	choices = map[string]int{
		"rock"		:	1
		"paper"		:	2
		"scissors"	:	3
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

