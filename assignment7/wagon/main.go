package main

import ( "./wagon"; "./screen" )

func main() {
	w := wagon.NewWagon()
	s := screen.NewScreen(24, 80)



	w.AddToHead(wagon.NewWheel("A", 5, 5))
	w.AddToHead(wagon.NewWheel("B", 5, 6))
	start := 'C'
	w.AddToScreen(s)
	s.Print()
	var c string
	for {
		c, _ = screen.ReadChar()
		switch c{

			case "U":
				w.Move(w.Tail, wagon.UP)
			case "D":
				w.Move(w.Tail, wagon.DOWN)
			case "L":
				w.Move(w.Tail, wagon.LEFT)
			case "R":
				w.Move(w.Tail, wagon.RIGHT)
			case "u":
				w.Move(w.Head, wagon.UP)
			case "d":
				w.Move(w.Head, wagon.DOWN)
			case "l":
				w.Move(w.Head, wagon.LEFT)
			case "r":
				w.Move(w.Head, wagon.RIGHT)
			case "q":
				break
			case "a":
				w.Add(w.Head, wagon.NewWheel((string)(start),0, 0))
				start+=1
			case "A":
				w.Add(w.Tail, wagon.NewWheel((string)(start), 24, 80))
				start +=1
		}

	}
}