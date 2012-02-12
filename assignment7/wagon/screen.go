/*
A package to represent a grid-based screen of certain integer height and width.
*/
package screen

import ( "os"; "exec"; "strconv"; "syscall"; "unsafe" )

/*
A Screen Struct to represent a screen with height and width.

Variables:
	height - the number of rows of the screen
	width - the number of columns of the screen
	buffer - the cahracter buffer to print to the screen
	on_screen - the cahracter buffer representing what is on the screen
	DefaultChar - the default character to draw to the screen if the buffer at that position is empty
*/
type Screen struct {
	height int
	width int
	buffer [][]string
	on_screen [][]string
	DefaultChar string
}

/*
Create a new screen.
Uses the terminal's dimensions for the screen's dimensions.

Returns:
	*Screen - a pointer to the screen representation
*/
func NewScreen() *Screen {
	s := new(Screen)
	s.height , s.width = screen_dimesnions()
	s.empty_buff()
	s.empty_screen()
	s.DefaultChar = " "
	return s
}

/*
Add a character to the screen at the specified row and column.

Method for:
	*Screen - a pointer to a screen struct

Parameters:
	chars - the characters to add to the screen (in string form)
	x - the x coordinate (column) to start adding the character to
	y - the y coordinate (row) to add the cahracters to

Returns:
	bool - true if successful, false if attempting to write out-of-bounds
*/
func (s *Screen) Add(chars string, x, y int) bool {
	if y > s.height { return false }
	for i, ch := range chars {
		col := x + i
		if col > s.width { return false }
		s.buffer[y][col] = (string)(ch)
	}
	return true
}

/*
Hides the cursor from the screen.
Always call ShowCursor() when you are done.
*/
func HideCursor() {
	os.Stdout.Write(([]byte)("\033[?25l"))
	os.Stdout.Sync()
}

/*
Shows the cursor on the screen.
*/
func ShowCursor() {
	os.Stdout.Write(([]byte)("\033[?25h"))
	os.Stdout.Sync()
}

/*
Print all of the characters from the screen buffer to standard out.

Method for:
	*Screen - a pointer to a screen struct
	
Post:
	s.buffer - the screen's buffer is reset
	s.on_screen - the on screen buffer matches what is on screen
*/
func (s *Screen) Print() {
	for r, line_arr := range s.buffer {
		for c, char := range line_arr {
			if char != s.on_screen[r][c] {
				if char == "" { char = s.DefaultChar }
				s.on_screen[r][c] = char
				result := "\033[" + strconv.Itoa(r) + ";" + strconv.Itoa(c) + "H" + char
				os.Stdout.Write(([]byte)(result))
				os.Stdout.Sync()
			}
		}
	}
	s.empty_buff()
}

/*
Clear all of the characters from the screen buffer and standard out.

Method for:
	*Screen - a pointer to a screen struct
	
Post:
	s.buffer - the screen's buffer is reset
	s.on_screen - the on screen buffer is reset
*/
func (s *Screen) Clear() {
	s.empty_buff()
	s.empty_screen()
	for r, line_arr := range s.on_screen {
		for c, char := range line_arr {
			if char == "" { char = s.DefaultChar }
			result := "\033[" + strconv.Itoa(r) + ";" + strconv.Itoa(c) + "H" + char
			os.Stdout.Write(([]byte)(result))
			os.Stdout.Sync()
		}
	}
}

/*
Change the dimensions of the screen.

Method for:
	*Screen - a pointer to a screen struct

Parameters:
	height - the new integer height (rows) of the screen
	width - the new integer width (columns) of the screen

Post:
	s.buffer - the screen's buffer is reset
	s.on_screen - the on-screen buffer is reset
*/
func (s *Screen) UpdateScreenSize() {
	row, col := screen_dimesnions()
	if row != s.height || col != s.width {
		s.height = row
		s.width = col
		s.empty_buff()
		s.empty_screen()
		s.Clear()
	}
}

/*
Get the dimensions of the screen.

Method for:
	*Screen - a pointer to a screen struct

Returns:
	height - the integer height (rows) of the screen
	width - the integer width (columns) of the screen
*/
func (s *Screen) GetDimensions() (int, int) {
	return s.height, s.width
}

/*
Empty the screen's buffer.
*/
func (s *Screen) empty_buff() {
	s.buffer = make([][]string, s.height+1)
	for i, _ := range s.buffer { s.buffer[i] = make([]string, s.width+1) }
}

/*
Empty the screen.
*/
func (s *Screen) empty_screen() {
	s.on_screen = make([][]string, s.height+1)
	for i, _ := range s.on_screen { s.on_screen[i] = make([]string, s.width+1) }
}

/*
Set up a command to use os standard in/out/err
*/
func set_exec(c *exec.Cmd) {
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
}

type winsize struct { 
    ws_row, ws_col uint16 
    ws_xpixel, ws_ypixel uint16 
}

/*
Get the dimensions of the standard out screen.

Returns:
	int - the number of rows on the screen
	int - the number of columns on the screen\
*/
func screen_dimesnions() (int, int) {
	ws := winsize{} 
    syscall.Syscall(syscall.SYS_IOCTL, 
        uintptr(0), uintptr(syscall.TIOCGWINSZ), 
        uintptr(unsafe.Pointer(&ws))) 
	return (int)(ws.ws_row), (int)(ws.ws_col)
}

/*
Reset the standard in and out screens from raw mode.

Pre:
	MakeRaw() is called, the screen is set to raw mode.
*/
var ResetRaw func()

/*
Set the standard in and out screens to raw mode.
Always call ResetRaw() when you are done.

Post:
	ResetRaw() will reset standard in and out to pre-raw settings
*/
func MakeRaw() {
	c := exec.Command("stty", "raw")
	set_exec(c)
	c.Run()

	ResetRaw = func() {
		c := exec.Command("stty", "-raw")
		set_exec(c)
		c.Run()
	}
}

/*
Read one character from standard input.

Returns:
	string - the character read from standard in
	os.Error - an error if the read went wrong, nil if successful
*/
func ReadChar() (string, os.Error) {
	one_char := " "
	char_byte := ([]byte)(one_char)
	_, err := os.Stdin.Read(char_byte)
	if err != nil { return "", err }
	result_char := (string)(char_byte)
	return result_char, nil
}