/*
A package to represent a collection of wheels as a wagon.
*/
package wagon

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

HasL
	Head - the beginning of the wagon train
	Tail - the end of the wagon train
*/
type Wagon struct {
	Head *Wheel
	Tail *Wheel
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
	w = new(Wheel)
	w.Value = value
	w.X = x
	w.Y = y
	return w
}

/*
Create a new Wheel.

Parameters:
	value - the value of the wheel

Returns:
	*Wheel - a pointer to a wheel with the given value
*/
func NewWheel(value string) *Wheel {
	w = new(Wheel)
	w.Value = value
	return w
}

/*
Create a new Wagon train.

Returns:
	*Wagon - a pointer to a fresh new wagon train.
*/
func NewWagon() *Wagon {
	w = new(Wagon)
	return w
}

/*
Add a wheel to the wagon train.

Method for:
	*Wagon - a pointer to the wagon to add a wheel to

Parameters:
	wheel - a pointer to the wheel to add to the wagon
*/
func (w *Wagon) Add(wheel *Wheel) {
	if w.Tail == nil {
		w.Head = wheel
		w.Tail = wheel
	} else {
		w.Tail.Next = wheel
		wheel.Prev = w.Tail
		w.Tail = wheel
	}
}