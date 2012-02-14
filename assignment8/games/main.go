package main

import ( "./games"; "./rps" )
//import "./ttt"

func main() {

	game := rps.NewGame()

	v1 := games.NewView()
	v2 := games.NewView()
	ref := games.NewReferee(game, v1, v2)

	go v1.Loop()
	go v2.Loop()

	ref.Loop()

}