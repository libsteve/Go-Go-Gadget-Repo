/*
A package to represent a collection of wheels as a wagon.
*/
package wagon

/*
An interfaces for anything that the wagon is going to write to.
*/
type Screen interface{
	/*
	Add the character to the string at the goven coordinates.

	Parameters:
		char - the character to add
		x - the x position to add to (the column)
		y - the y position to add to (the row)

	Returns:
		bool - true if successful, false otherwise
	*/
	Add(char string, x, y int) bool

	/*
	Get the dimensions of the screen.

	Returns:
		height - the integer height (rows) of the screen
		width - the integer width (columns) of the screen
	*/
	GetDimensions() (int, int) 
}

/*
Constant direction values.
*/
const (
	UP		= 0
	DOWN	= 1
	LEFT	= 2
	RIGHT	= 3
	HEAD 	= 4
	TAIL 	= 5
)

/*
The Wheel Struct.

Has:
	Value - represent the wheel
	X - the x position
	Y - the y position
	Next - the next wheel in the chain
	Prev - the previous wheel in the chain
*/
type Wheel struct {
	Value string
	X int
	Y int
	Next *Wheel
	Prev *Wheel
}

/*
The Wagon Struct.

Has:
	Head - the beginning of the wagon train
	Tail - the end of the wagon train
	height_bound - the maximum height of the wagon
	width_bound - the maximum width of the wagon
*/
type Wagon struct {
	Head *Wheel
	Tail *Wheel
	S Screen
}

/*
Create a new Wheel.

Parameters:
	value - the value of the wheel
	x - the x position
	y - the y position

Returns:
	*Wheel - a pointer to a wheel with the given values
*/
func NewWheel(value string, x int, y int) *Wheel {
	w := new(Wheel)
	w.Value = value
	w.X = x
	w.Y = y
	return w
}

/*
Create a new Wagon train.

Parameters:
	s - a pointer to some Screen interface

Returns:
	*Wagon - a pointer to a fresh new wagon train.
*/
func NewWagon(s Screen) *Wagon {
	w := new(Wagon)
	w.S = s
	return w
}

/*
Add a wheel to the head or tail of the wagon train.

Method for:
	*Wagon - a pointer to the wagon to add a wheel to

Parameters:
	dest - either HEAD or TAIL, depending on where the wheel is to be added
	wheel - a pointer to the wheel to add to the wagon
*/
func (w *Wagon) Add(dest int, wheel *Wheel) {
	if w.Tail == nil && (dest == HEAD || dest == TAIL) {
		w.Head = wheel
		w.Tail = wheel
	}
	if dest == HEAD {
		w.Head.Prev = wheel
		wheel.Next = w.Head
		w.Head = wheel
	} else if dest == TAIL {
		w.Tail.Next = wheel
		wheel.Prev = w.Tail
		w.Tail = wheel
	}

	
}

/*
Write the wagon to the given screen.

Post:
	The wagon data is given to the screen
*/
func (w *Wagon) AddToScreen() {
	var current *Wheel

	current = w.Head
	for {
		w.S.Add(current.Value, current.X, current.Y)
		if current.Next == nil { break }
		current = current.Next
	}
}

/*
Move the wagon either forward or backward.

Parameters:
	start - the start wheel, either HEAD or TAIL, for the head or tail of the wagon
	direction - the direction to move the start wheel. either UP, DOWN, LEFT, or RIGHT

Post:
	all the wheels within the wagon will move into the previous wheel's place
*/
func (w *Wagon) Move(start int, direction int) {
	var final *Wheel
	var current *Wheel
	var next *Wheel
	var next_wheel func() *Wheel
	var result_x, result_y int

	// set up variables
	switch start {
	case HEAD:
		final = w.Head
		result_x, result_y = dest_result(final, direction)
		current = w.Tail
		next_wheel = func() *Wheel { return current.Prev }
	case TAIL:
		final = w.Tail
		result_x, result_y = dest_result(final, direction)
		current = w.Head
		next_wheel = func() *Wheel { return current.Next }
	default:
		// do nothing
		return
	}

	// check bounds
	max_height, max_width := w.S.GetDimensions()
	if (result_x > max_width || result_y > max_height) || (result_x < 1 || result_y < 1) {
		return
	}

	// move the values throughout the linked-list
	for current != final {
		next = next_wheel()
		current.X, current.Y = next.X, next.Y
		current = next
	}

	final.X, final.Y = result_x, result_y
	
}

/*
Get the resulting x, y coordinates from the specified movement of the given wheel
*/
func dest_result(w *Wheel, direction int) (int, int) {
	x, y := w.X, w.Y
	switch direction {
	case UP:
		y -= 1
	case DOWN:
		y += 1
	case LEFT:
		x -= 1
	case RIGHT:
		x += 1
	}
	return x, y
}