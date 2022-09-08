// Package method value of expression
package method

// Method interface for script and golang,
// NOTE: args []interface{} lifecycle only in the method,
// DO NOT return the args or slice of args
type Method func(args []interface{}) (interface{}, error)
