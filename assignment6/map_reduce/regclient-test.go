package main

import (
	"./iregistry"
	"./regclient"
	"fmt"
	"rpc"
	"log"
)

func main() {
	nakedclient, err := rpc.DialHTTP("tcp", "localhost:"+(string)(registry.Port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	client := regclient.NewClientWrapper(nakedclient)

	_, err = client.Bind("banana-key", "banana")
	if err == nil {
		fmt.Println("banana-key bound")
	} else {
		fmt.Println("Bind Failure: " + err.String())
	}

	var value interface{}
	value, err = client.Lookup("banana-key", value.(string))
	if err == nil {
		fmt.Println("banana-key: bound to: " + value.(string))
	} else {
		fmt.Println("Lookup Failure: " + err.String())
	}

	_, err = client.Remove("banana-key", 0)
	if err == nil {
		fmt.Println("banana-key removed")
	} else {
		fmt.Println("Remove Failure: " + err.String())
	}

	err = client.Close()
	if err == nil {
		fmt.Println("successful close")
	} else {
		fmt.Println("Close FailureL " + err.String())
	}
}