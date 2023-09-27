// Code generated by "string-enumer -t ConnectionType -o ./generated.go ."; DO NOT EDIT.
package connectiontype

// validConnectionTypeValues contains a map of all valid ConnectionType values for easy lookup
var validConnectionTypeValues = map[ConnectionType]struct{}{
	Shell:     {},
	Dynamic:   {},
	Kube:      {},
	Web:       {},
	Db:        {},
	Ssh:       {},
	Rdp:       {},
	SqlServer: {},
}

// Valid validates if a value is a valid ConnectionType
func (v ConnectionType) Valid() bool {
	_, ok := validConnectionTypeValues[v]
	return ok
}

// ConnectionTypeValues returns a list of all (valid) ConnectionType values
func ConnectionTypeValues() []ConnectionType {
	return []ConnectionType{
		Shell,
		Dynamic,
		Kube,
		Web,
		Db,
		Ssh,
		Rdp,
		SqlServer,
	}
}
