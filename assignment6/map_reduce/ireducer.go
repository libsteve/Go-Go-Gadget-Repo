/*

The package reducer implements the proxy for remote procedure call access
over HTTP to execute a Job.

*/
package reducer

import (
	"mapper/mapper"
	"os"
)

// default service name
const Name = "Reducer"

type (
	// client-side view of the service.
	Reducer interface {
		// reserve, return status.
		Reserve() (bool, os.Error)

		// run a job of many tasks, return a result.
		Run(job Job) (interface{}, os.Error)

		// release.
		Release() os.Error

		// disconnect from the service.
		Close() os.Error
	}

	// to be implemented for a specific application.
	Job interface {
		// return all parallel tasks.
		Tasks() ([]mapper.Task, os.Error)

		// merge a task result into the job result.
		Reduce(result interface{})

		// retrieve the job result.
		Result() interface{}
	}
)