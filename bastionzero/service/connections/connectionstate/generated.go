// Code generated by "string-enumer -t ConnectionState -o ./generated.go ."; DO NOT EDIT.
package connectionstate

// validConnectionStateValues contains a map of all valid ConnectionState values for easy lookup
var validConnectionStateValues = map[ConnectionState]struct{}{
	Open:    {},
	Closed:  {},
	Error:   {},
	Pending: {},
}

// Valid validates if a value is a valid ConnectionState
func (v ConnectionState) Valid() bool {
	_, ok := validConnectionStateValues[v]
	return ok
}

// ConnectionStateValues returns a list of all (valid) ConnectionState values
func ConnectionStateValues() []ConnectionState {
	return []ConnectionState{
		Open,
		Closed,
		Error,
		Pending,
	}
}
