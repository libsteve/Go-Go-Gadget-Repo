package main

import (
	"os"
	"rpc"
	"net"
	"log"
	"http"
	"./registry"
)

type Reg struct {
	Map map[string]interface{}
}

func NewReg() *Reg {
	reg := new(Reg)
	reg.Map = make(map[string]interface{})
	return reg
}

func (reg *Reg) Bind(args registry.Bind, handle *int) os.Error {
	reg.Map[args.Key] = args.Data
	*handle = 0
	return nil
}

func (reg *Reg) Lookup(key string, pointer *interface{}) os.Error {
	result, ok := reg.Map[key]
	*pointer = result
	if !ok { return os.NewError(""); } 
	return nil
}

func (reg *Reg) Remove(args registry.Remove, number *int) os.Error {
	var generic interface{}
	reg.Map[args.Key] = generic, false
	*number = 1
	return nil
}

func main() {
	reg := NewReg()
	rpc.RegisterName(registry.Name, reg)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9901")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}