package wagon

type Wheel struct {
	Value string
	X int
	Y int
	Next Wheel
	Prev Wheel
}

type Wagon struct {
	Head Wheel
	Tail Wheel
}

