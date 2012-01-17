package dag

import(
	"os"
	"fmt"
)

/**
 * Represents an edge on a DAG
 */
type Edge_struct struct{
}

/**
 * Perform an action.
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
 * Represent as a string.
 */
func(edge Edge_struct) String() string {
	return ""
}

func MakeEdge() Edge_struct {
	return Edge_struct{}
}
