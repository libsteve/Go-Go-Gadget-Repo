package dag

import (
	"os"
)

/*
  Representst the dag
 */
type Dag_struct struct{
	edgeMap map[string] Edge
	connections [][]string
	visited []string
}

/*
  Makes a pointer to a new dag
 */
func MakeDag() *Dag_struct{
	news := new(Dag_struct)
	news.edgeMap = make(map[string]  Edge)
	news.connections = make([][]string, 0)
	return news
}
/*
  Adds an edge to the graph
  
  Parameters: 
  		targets - the list of targets
		sources - the list of sources to connect to the targets
		edge    - the edge you want to use to connect the target to sources

  Returns:
    	An Error if adding the target fails or nil if the add is successful
 */
func (d *Dag_struct) Add(targets, sources []string, edge Edge) os.Error{
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

/*
  Applys the action to the target using a depth first search
  
  Parameters:
  		string - the string representing your target
  
  Returns:
  		An error if applying the action fails
 */
func (d *Dag_struct) Apply(target string) os.Error{
	d.visited = append(d.visited, target)
	for _, v := range neighbors(d, target){
		if !visited(d,v){
			d.Apply(v)
		}
	}
	if( d.edgeMap[target] != nil){
		return d.edgeMap[target].Action(target, neighbors(d, target))
	}
	return nil
}

/*
  A string representiation of the dags
  
  Returns:
  		[(t1, s1), (t1, s2), (t2, s3), ...(tx, sy)]
 */
func (d *Dag_struct) String() string{
	str := "["
	for _, c := range d.connections{
		str += "(" + c[0] + ", " + c[1] + "), "
	}
	return str[0:len(str)- 2] + "]"
}

/*
  Have we been to this node before?

  Parameters:
  		d      - the dag struct 
		target - the target you are seeing if you have been to before
  
  Returns:
  		true if you have visited the node before
		false if not.
 */
func visited(d *Dag_struct , target string) bool{
	for _, prevVis := range d.visited{
		if prevVis == target{
			return true
		}
	}
	return false

}

/*
  Finds all of the neighbors of a given target

  Parameters:
  		d 	   - the Dag
		target - the target you are trying to find the neighbors of
  
  Returns:
  		an array of strings that contain the neighbors to the target
 */
func neighbors(d *Dag_struct, target string) []string{
	var vertices[]string
	for _, c := range d.connections {
		if c[0] == target{
			vertices = append(vertices, c[1])
		}
	}
	return vertices
}

