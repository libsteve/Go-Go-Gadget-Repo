package registry

func NewBind(key string, value interface{}) *Bind {
	bind := new(Bind)
	bind.Key = key
	bind.Data = value.([]byte)
	return bind
}

func NewRemove(key string, handle int) *Remove {
	remove := new(Remove)
	remove.Key = key
	remove.Handle = handle
	return remove
}