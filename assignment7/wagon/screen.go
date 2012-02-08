/*
A package to represent a grid-based screen of certain integer height and width.
*/
package screen

import ( "os"; "strconv"; "strings"; "exec" )
//import "terminal"

/*
A Screen Struct to represent a screen with height and width.

Variables:
	Height - the number of rows of the screen
	Width - the number of columns of the screen
	Buffer - the cahracter buffer to print to the screen
	DefaultChar - the default character to draw to the screen if the buffer at that position is empty
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
	s.EmptyBuffer()
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
	for r, line_arr := range s.Buffer {
		for c, char := range line_arr {
			if char == "" { char = s.DefaultChar }
			os.Stdout.Write(([]byte)("\033["+strconv.Itoa(r)+";"+strconv.Itoa(c)+"H"+char))
			os.Stdout.Sync()		
		}
	}
	s.EmptyBuffer()
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
	s.EmptyBuffer()
}

func (s *Screen) EmptyBuffer() {
	s.Buffer = make([][]string, s.Height)
	for i, _ := range s.Buffer { s.Buffer[i] = make([]string, s.Width) }
}

/*
Get the dimensions of the standard out screen.

Returns:
	int - the number of rows on the screen
	int - the number of columns on the screen
	os.Error - the error if the function failed, nil if successful
*/
func GetScreenDimensions() (int, int, os.Error) {
	c := exec.Command("stty", "size")
	raw_out, err := c.Output()
	if err == nil {
		out := (string)(raw_out)
		nums := strings.Split(out, " ")
		row, err1 := strconv.Atoi(nums[0])
		col, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil { return row, col, nil }
		return 0, 0, os.NewError("Convertion Error")
	}
	return 0, 0, err
}

/*
Prepare the standard in and out screens for raw mode.

Returns:
	func() os.Error - a function used for restoring the screens to their pre-raw mode
	os.Error - the error if the function failed, nil if successful
*/
func PrepScreenForRaw() (func() os.Error, os.Error) {
	c := exec.Command("stty", "-g")
	raw_out, err := c.Output()
	if err == nil {
		out := (string)(raw_out)
		reset_func := func() os.Error {
			c = exec.Command("stty", out)
			return c.Run()
		}
		return reset_func, nil
	}
	return func() os.Error { return err }, err
}