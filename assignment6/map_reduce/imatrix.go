/*

The package matrix implements the proxy for remote procedure call access
over HTTP to two-dimensional matrices with string names and float64 elements.

*/
package matrix

import "os"

// default service name
const Name = "Matrix"

// client-side view of the service.
type Matrix interface {

	// return the dimensions of an existing matrix.
	Dim(name string) (rows, cols int, err os.Error)

	// create a (unique) matrix with (positive) dimensions.
	Make(name string, rows, cols int) os.Error

	// delete an existing matrix.
	Remove(name string) os.Error

	// get an element value of an existing matrix.
	Get(name string, i, j int) (value float64, err os.Error)

	// set an element value of an existing matrix.
	Set(name string, i, j int, value float64) os.Error
	
	// disconnect from the service.
	Close() os.Error
}

// service's arguments (have to be public).
type (
	Make struct {
		Name       string
		Rows, Cols int
	}
	Get struct {
		Name string
		I, J int
	}
	Set struct {
		Name  string
		I, J  int
		Value float64
	}
)