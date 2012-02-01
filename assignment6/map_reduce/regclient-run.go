package main

import (
	"./regclient"
	"./parser"
	"fmt"
	"rpc"
	"log"
	"os"
)

func main() {
	nakedclient, err := rpc.DialHTTP("tcp", "localhost:9901")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	client := regclient.NewClientWrapper(nakedclient)

	cmds := parser.NewCommands()

	bind := func(args []string) os.Error {
		if len(args) == 2 {
			key := args[0]
			value := args[1]
			_, err = client.Bind(key, ([]byte)(value))
			if err == nil {
				fmt.Println("Successful Binding for KEY: "+key+" VALUE: "+value)
			} else {
				log.Fatal("Bind Failure: " + err.String())
			}
			return nil;
		}
		return os.NewError("Invalid Arguments")
	}

	cmds.AddInputCommand("bind", bind)

	lookup := func(args []string) os.Error {
		if len(args) == 1 && args[0] != "" {
			var value interface{}
			key := args[0]
			value, err = client.Lookup(key, value)
			result := (string)(value.([]byte))
			if err == nil {
				if result != "" {
					fmt.Println("KEY: "+key+" Bound to VALUE: "+result)
				} else {
					fmt.Println("KEY: "+key+" Not Found")
				}
			} else {
				log.Fatal("Lookup Failure: " + err.String())
			}
			return nil
		}
		return os.NewError("Invalid Arguments")
	}

	cmds.AddInputCommand("lookup", lookup)

	remove := func(args []string) os.Error {
		if len(args) == 1 && args[0] != "" {
			key := args[0]
			_, err= client.Remove(key, 0)
			if err == nil {
				fmt.Println("Removed KEY: "+key)
			} else {
				log.Fatal("Remove Failure: " + err.String())
			}
			return nil
		}
		return os.NewError("Invalid Arguments")
	}

	cmds.AddInputCommand("remove", remove)

	close_func := func() os.Error {
		err = client.Close()
		if err == nil {
			fmt.Println("successful close")
		} else {
			log.Fatal("Close Failure: " + err.String())
		}
		return os.NewError("CLOSE_PROGRAM")
	}

	cmds.AddCommand("close", close_func)

	// add read loop here

}