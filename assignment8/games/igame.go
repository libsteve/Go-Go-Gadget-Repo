package games

type Igame interface {
	
	CheckMoveValid(move string) bool
	
	MakeMove(player int, move string)

	Finished() bool

	Winner() int

}