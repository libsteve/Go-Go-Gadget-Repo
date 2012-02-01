package main

import (
	"./regclient"
	"./parser"
	"bufio"
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
		if len(args) == 2 && args[0] != "" && args[1] != "" {
			key := args[0]
			value := args[1]
			_, err = client.Bind(key, ([]byte)(value))
			if err == nil {
				fmt.Println("Successful Binding for KEY: "+key+" VALUE: "+value)
			} else {
				log.Println("Bind Failure: " + err.String())
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
			if err == nil {
				result := (string)(value.([]byte))
				fmt.Println("KEY: "+key+" Bound to VALUE: "+result)
			} else {
				log.Println("Lookup Failure: " + err.String())
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
				log.Println("Remove Failure: " + err.String())
			}
			return nil
		}
		return os.NewError("Invalid Arguments")
	}

	cmds.AddInputCommand("remove", remove)

	close_str := "CLOSE_PROGRAM"

	close_func := func() os.Error {
		err = client.Close()
		if err == nil {
			fmt.Println("successful close")
		} else {
			log.Fatal("Close Failure: " + err.String())
		}
		return os.NewError(close_str)
	}

	cmds.AddCommand("close", close_func)

	help := func() os.Error {
		fmt.Println("Command:\t\tDescription:")
		fmt.Println("")
		fmt.Println("help\t\t\tdisplay a list of commands")
		fmt.Println("close\t\t\tclose the scession")
		fmt.Println("bind: KEY, VALUE\tbind the KEY to the VALUE")
		fmt.Println("lookup: KEY\t\tfind the VALUE for th KEY")
		fmt.Println("remove: KEY\t\tremove the KEY and the VALUE")
		return nil
	}

	cmds.AddCommand("help", help)

	sin := bufio.NewReader(os.Stdin)
	var line string
	var error os.Error
	for error != os.EOF {
		fmt.Println("")
		line, error = sin.ReadString('\n')
		e := cmds.Parseln(line)
		if e != nil {
			if e.String() == close_str { 
				return
			} else {
				log.Println(e.String())
			}
		}
	}

}













