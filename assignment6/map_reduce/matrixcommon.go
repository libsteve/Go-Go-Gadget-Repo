package matrix

type Matrix_struct  struct {
	Name string
	Rows, Cols int
	Matrix [][]float64
}

func NewMatrix(name string, rows, cols int) *Matrix_struct{
	news := new(Matrix_struct)
	news.Name = name
	news.Rows = rows
	news.Cols = cols
	news.Matrix = make([][]float64, rows)
	for i, _ := range news.Matrix {
		news.Matrix[i] = make([]float64, cols)
	}
	return news
}

func NewMake(name string, rows, cols int) *Make{
	news := new(Make)
	news.Name = name
	news.Rows = rows
	news.Cols = cols
	return news
}

func NewGet(name string, i, j int) *Get{
	news := new(Get)
	news.Name = name
	news.I = i
	news.J = j
	return news
}

func NewSet(name string, i, j int, value float64) *Set{
	news := new(Set)
	news.Name = name
	news.I = i
	news.J = j
	news.Value = value
	return news
}