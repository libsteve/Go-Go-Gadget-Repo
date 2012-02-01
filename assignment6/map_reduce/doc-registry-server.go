/*
Runs the registry server.
Is the registry server.
*/
package doc-registry-server

/*
Create a new registry
*/
func NewReg() *Reg {}

/*
Bind a key to a value

Parameters:
	args - a bind struct with a key and data
	handle - a pointer to an int. this will be the return 
			 value for the handle that the key is assigned to

Returns:
	nil
*/
func (reg *Reg) Bind(args registry.Bind, handle *int) os.Error {}

/*
Find the value for a key

Parameters:
	key - the key string to look for
	pointer - a pointer to anything. this will be the returned data for the key

Returns:
	os.Error if the key is not found, nil otherwise
*/
func (reg *Reg) Lookup(key string, pointer *interface{}) os.Error {}

/*
Remove a key and value

Parameters:
	args - a remove struct with a key and number to remove
	number - a pointer to an int. this will be the return 
			 the amount of entries removed

Returns:
	os.Error if the key is not found, nil otherwise
*/
func (reg *Reg) Remove(args registry.Remove, number *int) os.Error {}