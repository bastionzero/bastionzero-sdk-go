// Code generated by "string-enumer -t RoleType -o ./generated.go ."; DO NOT EDIT.
package roletype

// validRoleTypeValues contains a map of all valid RoleType values for easy lookup
var validRoleTypeValues = map[RoleType]struct{}{
	User:  {},
	Admin: {},
}

// Valid validates if a value is a valid RoleType
func (v RoleType) Valid() bool {
	_, ok := validRoleTypeValues[v]
	return ok
}

// RoleTypeValues returns a list of all (valid) RoleType values
func RoleTypeValues() []RoleType {
	return []RoleType{
		User,
		Admin,
	}
}