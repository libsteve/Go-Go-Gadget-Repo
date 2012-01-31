package regclient

import (
	"os"
	"rpc"
	"./registry"
)

type ClientWrapper struct {
	Client *rpc.Client
}

func NewClientWrapper(nakedclient *rpc.Client) *ClientWrapper {
	client := new(ClientWrapper)
	client.Client = nakedclient
	return client
}

// enter another (non-nil) value for a (non-empty) key
// and return a handle for removal.
func (client *ClientWrapper) Bind(key string, value interface{}) (int, os.Error) {
	bind := registry.NewBind(key, value)
	var handle int
	err := client.Client.Call("Registry.Bind", bind, &handle)
	return handle, err
}

// return all values for a (non-empty) key.
// pointer must be a slice of the values to be returned;
// a new (non-empty) slice will be returned.
func (client *ClientWrapper) Lookup(key string, pointer interface{}) (interface{}, os.Error) {
	err := client.Client.Call("Registry.Lookup", key, &pointer)
	return pointer, err
}

// remove a value for an existing (non-empty) key
// and the correct handle, or zero to remove all;
// return the positive number of values removed, or an error.
func (client *ClientWrapper) Remove(key string, handle int) (int, os.Error) {
	remove := registry.NewRemove(key, handle)
	var number int
	err := client.Client.Call("Registry.Remove", remove, &number)
	return number, err
}

// disconnect from the service.
func (client *ClientWrapper) Close() os.Error {
	return client.Client.Close()
}