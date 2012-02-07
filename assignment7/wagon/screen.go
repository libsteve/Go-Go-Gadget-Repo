/*
A package to represent a grid-based screen of certain integer height and width.
*/
package screen

import "fmt"
//import "terminal"

/*
A Screen Struct to represent a screen with height and width.
*/
type Screen struct {
	Height int
	Width int
	Buffer [][]string
	DefaultChar string
}

/*
Create a new screen.

Parameters:
	height - the integer height (rows) of the screen
	width - the integer width (columns) of the screen

Returns:
	*Screen - a pointer to the screen representation
*/
func NewScreen(height int, width int) *Screen {
	s := new(Screen)
	s.Height = height
	s.Width = width
	s.Buffer = make([][]string, height)
	for i, _ := range s.Buffer { s.Buffer[i] = make([]string, width) }
	s.DefaultChar = " "
	return s
}

/*
Create a new screen.
Uses the terminal's dimensions for the screen's dimensions.
Terminal only exists on the weekly build of GO.

Returns:
	*Screen - a pointer to the screen representation
*/
//func NewScreen() *Screen {
//	s = new(Screen)
//	s.Width, s.Height = terminal.GetSize(1) // fd = 1 is standard out
//	s.Buffer = new([s.Height][s.Width]string)
//	s.DefaultChar = " "
//	return s
//}

/*
Add a character to the screen at the specified x (row) and y (column) coordinates.

Method for:
	*Screen - a pointer to a screen struct

Parameters:
	char - the character to add to the screen
	x - the x coordinate (the row) to add the cahracter to
	y - the y coordinate (the column) to add the character to 

Returns:
	bool - true if the character was added, false otherwise

Pre:
	char - the character must be a single character. will return false otherwise.
*/
func (s *Screen) Add(char string, x int, y int) bool{
	if s.Buffer[x][y] == "" {
		if len(char) > 1 { return false }
		s.Buffer[x][y] = char
		return true
	}
	return false
}

/*
Print all of the characters from the screen buffer to standard out.

Method for:
	*Screen - a pointer to a screen struct
	
Post:
	s.Buffer - the screen's buffer is reset
*/
func (s *Screen) Print() {
	for _, line_arr := range s.Buffer {
		var line string
		for _, char := range line_arr {
			if char == "" { line += s.DefaultChar } else { line += char }
		}
		fmt.Println(line)
	}
	s.Buffer = make([][]string, s.Height)
	for i, _ := range s.Buffer { s.Buffer[i] = make([]string, s.Width) }
}

/*
Change the dimensions of the screen.

Method for:
	*Screen - a pointer to a screen struct

Parameters:
	height - the new integer height (rows) of the screen
	width - the new integer width (columns) of the screen

Post:
	s.Buffer - the screen's buffer is reset
*/
func (s *Screen) ChangeScreenSize(height int, width int) {
	s.Height = height
	s.Width = width
	s.Buffer = make([][]string, height)
	for i, _ := range s.Buffer { s.Buffer[i] = make([]string, width) }
}