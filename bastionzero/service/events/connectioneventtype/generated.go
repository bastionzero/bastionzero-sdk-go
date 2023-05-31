// Code generated by "string-enumer -t ConnectionEventType -o ./generated.go ."; DO NOT EDIT.
package connectioneventtype

// validConnectionEventTypeValues contains a map of all valid ConnectionEventType values for easy lookup
var validConnectionEventTypeValues = map[ConnectionEventType]struct{}{
	Created:          {},
	Closed:           {},
	ClientConnect:    {},
	ClientDisconnect: {},
	ShellConnect:     {},
	ShellDisconnect:  {},
}

// Valid validates if a value is a valid ConnectionEventType
func (v ConnectionEventType) Valid() bool {
	_, ok := validConnectionEventTypeValues[v]
	return ok
}

// ConnectionEventTypeValues returns a list of all (valid) ConnectionEventType values
func ConnectionEventTypeValues() []ConnectionEventType {
	return []ConnectionEventType{
		Created,
		Closed,
		ClientConnect,
		ClientDisconnect,
		ShellConnect,
		ShellDisconnect,
	}
}
