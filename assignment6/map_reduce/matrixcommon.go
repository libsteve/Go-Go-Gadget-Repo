package matrixcommon

type Matrix_struct  struct {
	Name string
	Rows, Cols int
	Matrix [][]float64
}

func New(name string, rows, cols int) *Matrix_struct{
	news := new(Matrix_struct)
	news.Name = name
	news.Rows = rows
	news.Cols = cols
	news.Matrix = make([][]float64, rows)
	for i, _ := range news.Matrix {
		m[i] = make([]int, cols)
	}
	return news
}