package dag

import(
	"os"
	"fmt"
	"exec"
	"strings"
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
	var error os.Error
	error = nil
	for _, command := range edge.Commands {
		cmdLine := strings.Split(command, " ")
		cmd := exec.Command(cmdLine[0], cmdLine[1:]...)
		output, error := cmd.CombinedOutput()
		fmt.Println(string(output))
		if error != nil { return error; }
	}
	return error
}

/**
 Represent as a string.

 Returns:
	Returns a string 
 */
func(edge Edge_struct) String() string {
	str := ""
	for _,command := range edge.Commands{
		str += command + "\n"
	}
	return str
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
