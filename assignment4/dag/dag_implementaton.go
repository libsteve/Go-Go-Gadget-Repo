package dag

import (
	"os"
)

type Dag_struct struct{
	edgeMap map[string] Edge
	connections [][]string

}

func MakeDag() *Dag_struct{
 	var news Dag_struct
	news.edgeMap = make(map[string]  Edge)
	news.connections = make([][]string, 0)
	return &news
}

func (d Dag_struct) Add(targets, sources []string, edge Edge) os.Error{
	if len(targets) == 0{
		return os.NewError("targets cannot be empty")
	}
	for _, t := range targets {
		d.edgeMap[t] = edge, true
		for _, s := range sources{
			d.connections = append(d.connections, []string{t,s})
		}
	}
	return nil
}

func (d Dag_struct) Apply(target string) os.Error{
	visited := make([]string, 0)
	return dfs(d, target, &visited)
}

func (d Dag_struct) String() string{
	str := ""
	for _, c := range d.connections{
		str += "(" + c[0] + ", " + c[1] + "), "
	}
	return str
}

func dfs(d Dag_struct, vertex string, visited *[]string) os.Error{
	*visited = append(*visited, vertex)
	for _, v := range adjacent(d, vertex){
		if !contains(visited, v){
			dfs(d, v, visited)
		}
	}
	return d.edgeMap[vertex].Action(vertex, adjacent(d, vertex))
}


func contains(visited* []string, visitor string) bool{
	for _, prevVis := range *visited{
		if prevVis == visitor{
			return true
		}
	}
	return false

}

func adjacent(d Dag_struct, vertex string) []string{
	var vertices[]string
	for _, c := range d.connections {
		if c[0] == vertex{
			vertices = append(vertices, c[1])
		}
	}
	return vertices
}

