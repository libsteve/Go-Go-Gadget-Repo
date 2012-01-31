package main

import(
	"rpc"
	"./imatrix"
	"os"
	"fmt"
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

func main (){
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	cw := NewClientWrapper(client)
	cw.Make("Awesome", 5, 10)
	r, c, _ := cw.Dim("Awesome")
	fmt.Printf("R: %d C: %d \n", r, c)
	count := 0

	for i := 0 ; i < 5; i++ {
		for j := 0; j <10 ; j++ {
			cw.Set("Awesome", i, j, float64(count))
			count++
		}
	}
	val, _ := cw.Get("Awesome", 2, 4)
	fmt.Printf("I: %d J: %d Val: %f \n", 2, 4 ,val )
	cw.Remove("Awesome")
}