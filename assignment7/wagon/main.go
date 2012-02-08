package main

import ( "./wagon"; "./screen" )

func main() {
	w := wagon.NewWagon()
	s := screen.NewScreen(24, 80)



	w.Add(wagon.NewWheel("A", 5, 5))
	w.AddToScreen(s)
	s.Print()
}