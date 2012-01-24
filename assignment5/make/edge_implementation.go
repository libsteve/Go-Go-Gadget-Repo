package edge

import(
	"os"
	"fmt"
)

/**
 Represents an edge on a DAG
 */
type Edge_struct struct{
	Commands []string
}

/**
 Perform an action.
 */
func(edge Edge_struct) Action(target string, sources []string) os.Error {
	result := target + ":"
	for _, source := range sources {
		result += " " + source
	}
	fmt.Println(result)
	return nil
}

/**
 Represent as a string.

 Returns:
	Returns a string 
 */
func(edge Edge_struct) String() string {
	return ""
}

/**
 Creates an Edge_struct

 Returns:
	A new Edge_struct
 */
func MakeEdge(c []string) *Edge_struct {
	news := new(Edge_struct)
	news.Commands = c
	return news
}
