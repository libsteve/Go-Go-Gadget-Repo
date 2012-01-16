package dag

import (
	"os"
)

type Dag_struct struct{

	edges []Edge_struct
	connections [][]string

}

func (d Dag_struct) Add(targets, sources []string, edge Edge_struct) os.Error{
	if len(targets) == 0{
		return os.NewError("targets cannot be empty")
	}
	for _, t := range targets {
		edge.targets = append(edge.targets, t)
		for _, s := range sources{
			d.connections = append(d.connections, []string{t,s})
		}
	}
	d.edges = append(d.edges, edge)
	return nil
}

func (d Dag_struct) Apply(target string) os.Error{
	return nil
}
