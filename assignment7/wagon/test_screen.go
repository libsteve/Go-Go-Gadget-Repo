package main

import "./screen"
import "exec"

func main() {

	screen.PrepScreenForRaw()
	raw := exec.Command("stty", "raw")
	err := raw.Run()
	if err != nil { println(err.String()) }

	s := screen.NewScreen(64, 82)
	for i := 0; i < 20; i++ {
		s.Add("A", 10+i, 10+i)
		s.Print()
		sleep := exec.Command("sleep", "2")
		sleep.Run()
	}
}