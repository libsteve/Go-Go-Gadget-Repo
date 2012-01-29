/*

The package registry implements the proxy for remote procedure call access
over HTTP to a map with string keys and byte slice values.

*/
package registry

import "os"

const (
	// default port (on localhost)
	Port = 9901

	// default service name
	Name = "Registry"
)

// client-side view of the service.
type Registry interface {

	// enter another (non-nil) value for a (non-empty) key
	// and return a handle for removal.
	Bind(key string, value interface{}) (int, os.Error)

	// return all values for a (non-empty) key.
	// pointer must be a slice of the values to be returned;
	// a new (non-empty) slice will be returned.
	Lookup(key string, pointer interface{}) (interface{}, os.Error)

	// remove a value for an existing (non-empty) key
	// and the correct handle, or zero to remove all;
	// return the positive number of values removed, or an error.
	Remove(key string, handle int) (int, os.Error)
	
	// disconnect from the service.
	Close() os.Error
}

// service's arguments (have to be public). 
type (
	Bind struct {
		Key  string
		Data []byte
	}
	Remove struct {
		Key    string
		Handle int
	}
)