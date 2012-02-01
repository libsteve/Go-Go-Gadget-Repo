package main

import(
	"rpc"
	"./imatrix"
	"./parser"
	"strconv"
	"os"
	"fmt"
	"bufio"
	"log"
)

type Client_wrapper struct{
	Client *rpc.Client
}

func NewClientWrapper(client *rpc.Client) *Client_wrapper{
	news := new(Client_wrapper)
	news.Client = client;
	return news
}

// return the dimensions of an existing matrix.
func (client *Client_wrapper) Dim(name string) (rows, cols int, err os.Error){
	var reply matrix.Matrix_struct
	err = client.Client.Call("Matrix_database.Dim", name, &reply)
	if err != nil {
		log.Fatal("Dimensions error:", err)
	}
	return reply.Rows, reply.Cols, nil
}

// create a (unique) matrix with (positive) dimensions.
func (client *Client_wrapper) Make(name string, rows, cols int) os.Error{
	var reply bool
	mk := matrix.NewMake(name, rows, cols)
	err := client.Client.Call("Matrix_database.Make", *mk, &reply)
	if err != nil {
		log.Fatal("Make error:", err)
	}
	return nil
}

// delete an existing matrix.
func (client *Client_wrapper) Remove(name string) os.Error{
	var reply bool
	err := client.Client.Call("Matrix_database.Remove", name, &reply)
	if err != nil {
		log.Fatal("Remove error:", err)
	}
	return nil
}

// get an element value of an existing matrix.
func (client *Client_wrapper) Get(name string, i, j int) (value float64, err os.Error){
	var reply float64
	get := matrix.NewGet(name, i, j)
	err = client.Client.Call("Matrix_database.Get", get, &reply)
	if err != nil {
		log.Fatal("Get error:", err)
	}
	return reply, nil
}

// set an element value of an existing matrix.
func (client *Client_wrapper) Set(name string, i, j int, value float64) os.Error{
	var reply bool
	set := matrix.NewSet(name, i, j, value)
	err := client.Client.Call("Matrix_database.Set", set, &reply)
	if err != nil {
		log.Fatal("Set error:", err)
	}
	return nil
}
	
// disconnect from the service.
func (client *Client_wrapper) Close() os.Error{
	err := client.Client.Close();
	if err != nil{
		log.Fatal("Close error", err)
	}
	return nil
}

func make_commands(cw *Client_wrapper) *parser.Commands{
	commands := parser.NewCommands()
	dim := func(input []string) os.Error{
		if len(input) != 1{
			return os.NewError("Invalid arguments")
		}
		r, c, err := cw.Dim(input[0])
		if err == nil{
			fmt.Printf("R: %d C: %d \n", r, c)
			return nil
		}
		return err
	}
	mak := func(input []string) os.Error{
		if (len(input) != 3){
			return os.NewError("Invalid arguments")
		}
		rows, _ := strconv.Atoi(input[1])
		cols, _ := strconv.Atoi(input[2])
		return cw.Make(input[0], rows, cols)
	}
	rm := func (input []string) os.Error{
		if (len(input) != 1){
			return os.NewError("Invalid arguments")
		}
		return cw.Remove(input[0])
	}

	get := func(input []string) os.Error{
		if len(input) != 3{
			return os.NewError("Invalid arguments")
		}
		i, _ := strconv.Atoi(input[1])
		j, _ := strconv.Atoi(input[2])
		value, err := cw.Get(input[0], i, j)
		if err == nil{
			fmt.Printf("I: %d J: %d Val: %f \n", i, j ,value )
		}
		return err
		
	}

	set := func(input []string) os.Error{
		if len(input) != 4{
			return os.NewError("Invalid arguments")
		}
		i, _ := strconv.Atoi(input[1])
		j, _ := strconv.Atoi(input[2])
		v, _ := strconv.Atof64(input[3])
		return cw.Set(input[0], i, j, v)
		
	}

	clo := func() os.Error{
		return cw.Close()
	}

	commands.AddInputCommand("Make", mak)
	commands.AddInputCommand("Dim", dim)
	commands.AddInputCommand("Remove", rm)
	commands.AddInputCommand("Get", get)
	commands.AddInputCommand("Set", set)
	commands.AddCommand("Close", clo)
	return commands
}

func main (){
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	cw := NewClientWrapper(client)
	commands := make_commands(cw)
	sin := bufio.NewReader(os.Stdin)
	var l string
	for err != os.EOF{
		l,err = sin.ReadString('\n')
		e := commands.Parseln(l)
		if (e != nil){
			fmt.Fprintln(os.Stderr, e.String())
		}
	}
}