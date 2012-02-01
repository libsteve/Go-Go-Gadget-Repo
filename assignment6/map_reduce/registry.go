package registry

/*
create a new bind struct
*/
func NewBind(key string, value interface{}) *Bind {
	bind := new(Bind)
	bind.Key = key
	bind.Data = value.([]byte)
	return bind
}

/*
create a new remove struct
*/
func NewRemove(key string, handle int) *Remove {
	remove := new(Remove)
	remove.Key = key
	remove.Handle = handle
	return remove
}