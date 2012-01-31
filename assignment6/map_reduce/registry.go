package registery

import (
	"os"
)

type Registry struct {
	
}

// enter another (non-nil) value for a (non-empty) key
// and return a handle for removal.
func (r *Registry) Bind(key string, value interface{}) (int, os.Error) {
	
}

// return all values for a (non-empty) key.
// pointer must be a slice of the values to be returned;
// a new (non-empty) slice will be returned.
func (r *Registry) Lookup(key string, pointer interface{}) (interface{}, os.Error) {
	
}

// remove a value for an existing (non-empty) key
// and the correct handle, or zero to remove all;
// return the positive number of values removed, or an error.
func (r *Registry) Remove(key string, handle int) (int, os.Error) {
	
}

// disconnect from the service.
func (r *Registry) Close() os.Error {
	
}