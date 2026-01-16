package safeconv

import "reflect"

// typeName returns the name of a type for error messages.
func typeName[T any]() string {
	var zero T
	return reflect.TypeOf(zero).Name()
}
