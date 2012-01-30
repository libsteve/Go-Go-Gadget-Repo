package matrix

type Matrix_struct{
	Name string
	Rows, Cols int
	Matrix [][]float64
}

type Matrix_database{
	Matrices []Matrix_struct
}
// return the dimensions of an existing matrix.
func (m *Matrix_database) Dim(name string) (rows, cols int, err os.Error){
	
}

// create a (unique) matrix with (positive) dimensions.
func (m *Matrix_database) Make(name string, rows, cols int) os.Error{
	
}

// delete an existing matrix.
func (m *Matrix_database) Remove(name string) os.Error{
	
}

// get an element value of an existing matrix.
func (m *Matrix_database) Get(name string, i, j int) (value float64, err os.Error){
	
}

// set an element value of an existing matrix.
func (m *Matrix_database) Set(name string, i, j int, value float64) os.Error{
	
}
	
// disconnect from the service.
func (m *Matrix_database) Close() os.Error{
	
}