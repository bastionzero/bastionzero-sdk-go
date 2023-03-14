// Code generated by "string-enumer -t SubjectType -o ./generated.go ."; DO NOT EDIT.
package subjecttype

// validSubjectTypeValues contains a map of all valid SubjectType values for easy lookup
var validSubjectTypeValues = map[SubjectType]struct{}{
	User:           {},
	ApiKey:         {},
	ServiceAccount: {},
}

// Valid validates if a value is a valid SubjectType
func (v SubjectType) Valid() bool {
	_, ok := validSubjectTypeValues[v]
	return ok
}

// SubjectTypeValues returns a list of all (valid) SubjectType values
func SubjectTypeValues() []SubjectType {
	return []SubjectType{
		User,
		ApiKey,
		ServiceAccount,
	}
}