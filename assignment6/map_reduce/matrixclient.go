package main

import(
	"rpc"
	"imatrix"
	"os"
	"matrixcommon"
)

// return the dimensions of an existing matrix.
func (m *rpc.Client) Dim(name string) (rows, cols int, err os.Error){

}

// create a (unique) matrix with (positive) dimensions.
func (m *rpc.Client) Make(name string, rows, cols int) os.Error{
	
}

// delete an existing matrix.
func (m *rpc.Client) Remove(name string) os.Error{
	
}

// get an element value of an existing matrix.
func (m *rpc.Client) Get(name string, i, j int) (value float64, err os.Error){
	
}

// set an element value of an existing matrix.
func (m *rpc.Client) Set(name string, i, j int, value float64) os.Error{
	
}
	
// disconnect from the service.
func (m *rpc.Client) Close() os.Error{
	
}