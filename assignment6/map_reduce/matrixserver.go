package main

import(
	"matrixcommon"
	"imatrix"
	"os"
)

/**
 * The Matrix database. Contains a map of string names associated with  
 * matrix structs
 */
type Matrix_database{
	Matrices map[string]Matrix_struct
}

func (m *Matrix_database) Dim(name string, matrix *Matrix_struct) os.Error{
	
}

func (m *Matrix_database) Make(make Make) os.Error{
}

func (m *Matrix_database) Remove(name string) os.Error{
	
}

func (m *Matrix_database) Get(get Get, val *float64) os.Error{
	
}

func (m *Matrix_database) Set(set Set) os.Error{
	
}