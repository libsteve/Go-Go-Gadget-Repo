package wagon

type Wheel struct {
	Value string
	X int
	Y int
	Next *Wheel
	Prev *Wheel
}

type Wagon struct {
	Head *Wheel
	Tail *Wheel
}

func NewWheel(value string) *Wheel {
	w = new(Wheel)
	w.Value = value
	return w
}

func NewWagon() *Wagon {
	w = new(Wagon)
	return w
}

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