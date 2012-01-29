/*

The package mapper implements the proxy for remote procedure call access
over HTTP to execute a Task.

*/
package mapper

import "os"

// default service name
const Name = "Mapper"

type (
	// client-side view of the service.
	Mapper interface {
		// reserve, return status.
		Reserve() (bool, os.Error)

		// run a task, return a result.
		Run(task Task) (interface{}, os.Error)

		// release.
		Release() os.Error

		// disconnect from the service.
		Close() os.Error
	}

	// to be implemented for a specific application.
	Task interface {
		Run() (interface{}, os.Error)
	}
)