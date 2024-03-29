// Code generated by "string-enumer -t VerbType -o ./generated.go ."; DO NOT EDIT.
package verbtype

// validVerbTypeValues contains a map of all valid VerbType values for easy lookup
var validVerbTypeValues = map[VerbType]struct{}{
	Shell:        {},
	FileTransfer: {},
	Tunnel:       {},
	RDP:          {},
	SQLServer:    {},
}

// Valid validates if a value is a valid VerbType
func (v VerbType) Valid() bool {
	_, ok := validVerbTypeValues[v]
	return ok
}

// VerbTypeValues returns a list of all (valid) VerbType values
func VerbTypeValues() []VerbType {
	return []VerbType{
		Shell,
		FileTransfer,
		Tunnel,
		RDP,
		SQLServer,
	}
}
