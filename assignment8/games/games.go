/*

Package games implements referees for the games of Rock Paper Scissors
and Tic Tac Toe with an active view connected to a terminal
and proxy views connected to a network.
The connection to any view is based on channels
so that it can be distributed using the netchan package.

*/
package games

import "os"

type (
	// sent to the view.
	Command int

	// describe a game's outcome.
	Outcome string
)

const (
	// enable user interface until view's human player
	// selects a move; no response.
	Enable = Command(iota)

	// respond with the view's human player's move. If the move is
	// illegal, Enable and Get will be sent again.
	Get

	// store other view's human player's move; no response.
	Set

	// display other view's human player's move; no response.
	Show

	// report game's outcome to the view's human player; no response.
	Done

	// neither human player wins.
	Draw = Outcome("draw")

	// view's human player wins.
	Win = Outcome("win")

	// view's human player loses.
	Lose = Outcome("lose")
)

type (
	// sent to a view; Set has one or more arguments, Done has one Outcome.
	Request struct {
		Command
		Args []string
	}

	// non-empty, received from a view in response to Get.
	Response []string

	// bi-directional connections between a view and a referee.
	View struct {
		Request  chan Request
		Response chan Response
	}

	// an active view has a run loop.
	Looper interface {
		Loop() os.Error
	}
)