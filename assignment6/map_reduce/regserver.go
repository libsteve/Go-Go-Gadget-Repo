/*
Runs the registry server.
Is the registry server.
*/
package main

import (
	"os"
	"rpc"
	"net"
	"log"
	"http"
	"strconv"
	"./registry"
)

/*
A registry struct
*/
type Reg struct {
	Map map[string]interface{}
}

/*
Create a new registry
*/
func NewReg() *Reg {
	reg := new(Reg)
	reg.Map = make(map[string]interface{})
	return reg
}

/*
Bind a key to a value

Parameters:
	args - a bind struct with a key and data
	handle - a pointer to an int. this will be the return 
			 value for the handle that the key is assigned to

Returns:
	nil
*/
func (reg *Reg) Bind(args registry.Bind, handle *int) os.Error {
	reg.Map[args.Key] = args.Data
	*handle = 0
	return nil
}

/*
Find the value for a key

Parameters:
	key - the key string to look for
	pointer - a pointer to anything. this will be the returned data for the key

Returns:
	os.Error if the key is not found, nil otherwise
*/
func (reg *Reg) Lookup(key string, pointer *interface{}) os.Error {
	result, ok := reg.Map[key]
	*pointer = result
	if !ok { return os.NewError("Key Not Found") } 
	return nil
}

/*
Remove a key and value

Parameters:
	args - a remove struct with a key and number to remove
	number - a pointer to an int. this will be the return 
			 the amount of entries removed

Returns:
	os.Error if the key is not found, nil otherwise
*/
func (reg *Reg) Remove(args registry.Remove, number *int) os.Error {
	var generic interface{}
	if _, ok := reg.Map[args.Key]; !ok { return os.NewError("Key Not Found") }
	reg.Map[args.Key] = generic, false
	*number = 1
	return nil
}

func main() {
	reg := NewReg()
	rpc.RegisterName(registry.Name, reg)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+strconv.Itoa(registry.Port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}