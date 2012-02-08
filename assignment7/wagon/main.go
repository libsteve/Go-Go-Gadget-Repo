package main

import ( "./wagon"; "./screen" )

func main() {
	w := wagon.NewWagon()
	s := screen.NewScreen()

	w.Add(wagon.NewWheel("A", 5, 5))
	w.AddToScreen(s.Add)
}