package main

import ( "./wagon"; "./screen"; "flag" )

func main() {
	s := screen.NewScreen()
	w := wagon.NewWagon(s)

	var d *string
	d = flag.String("b", " ", "set the background filler character")
	flag.Parse()
	if len(*d) == 1 { 
		s.DefaultChar = *d
	} else { 
		println("Only be one character allowed.")
		return 
	}

	screen.MakeRaw()
	screen.HideCursor()
	s.Clear()

	run(w, s)

	screen.ResetRaw()
	screen.ShowCursor()
}

func run(w *wagon.Wagon, s *screen.Screen) {
	HEAD, TAIL := wagon.HEAD, wagon.TAIL

	btm_y, btm_x := s.GetDimensions()

	start := 'a' - 1
	next_char := func() string {
		if start == 'z' { start = 'A' - 1 }
		if start == 'Z' { start = '0' - 1 }
		if start == '9' { start = 'a' - 1 }
		start += 1
		return (string)(start)
	}

	w.Add(HEAD, wagon.NewWheel(next_char(), 5, 5))
	w.Add(TAIL, wagon.NewWheel(next_char(), btm_x - 5, btm_y - 5))
	var c string
	main_loop: for {
		s.UpdateScreenSize()
		w.AddToScreen()
		s.Print()
		btm_y, btm_x = s.GetDimensions()
		c, _ = screen.ReadChar()
		switch c {
		case "U":
			w.Move(TAIL, wagon.UP)
		case "D":
			w.Move(TAIL, wagon.DOWN)
		case "L":
			w.Move(TAIL, wagon.LEFT)
		case "R":
			w.Move(TAIL, wagon.RIGHT)
		case "u":
			w.Move(HEAD, wagon.UP)
		case "d":
			w.Move(HEAD, wagon.DOWN)
		case "l":
			w.Move(HEAD, wagon.LEFT)
		case "r":
			w.Move(HEAD, wagon.RIGHT)
		case "q":
			break main_loop
		case "\033":
			next, _ := screen.ReadChar()
			if next == "[" {
				next, _ = screen.ReadChar()
				switch next {
				case "A":
					w.Move(HEAD, wagon.UP)
				case "B":
					w.Move(HEAD, wagon.DOWN)
				case "C":
					w.Move(HEAD, wagon.RIGHT)
				case "D":
					w.Move(HEAD, wagon.LEFT)
				}
			} else if next == "\033" { break main_loop }
		case "a":
			w.Add(HEAD, wagon.NewWheel(next_char(), 2, 2))
		case "A":
			w.Add(TAIL, wagon.NewWheel(next_char(), btm_x - 2, btm_y - 2))
		}
	}
}