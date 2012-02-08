package main

import "./screen"
import ( "exec"; "os" )

func main() {

	os.Exec("stty", []string{"raw"}, os.Environ() )

	row, col, err := screen.GetScreenDimensions()
	println(row)
	println(col)

	if err == nil {
		s := screen.NewScreen(row, col)
		reset, err := screen.PrepScreenForRaw()
		if err != nil {
			for i := 0; i < 20; i++ {
				s.Add("A", 10+i, 10+1)
				s.Print()
				sleep := exec.Command("sleep", "1")
				sleep.Run()
			}
			reset()
		}
	}
}