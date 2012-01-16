package dag

import(
	"os"
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
	return nil
}

/**
 * Represent as a string.
 */
func(edge Edge_struct) String() string {
	return ""
}
