/*

Package dag implements a directed acyclic graph
where edge actions can be applied depth-first, i.e.,
from the leaves to the root of a subtree.

*/
package dag

import "os"

type (
	// What a directed, acyclic graph must do.
	Dag interface {

		// Add an edge.
		//   targets cannot be empty,
		//   sources are empty if the targets are leaf nodes,
		//   all elements should not be blank.
		Add(targets, sources []string, edge Edge) os.Error

		// Apply the action, depth-first, to each edge
		// leading to a target.
		Apply(target string) os.Error

		// Represent as a string.
		String() string
	}

	// What an edge of the graph must do.
	Edge interface {

		// Perform an action.
		Action(target string, sources []string) os.Error

		// Represent as a string.
		String() string
  }
)
