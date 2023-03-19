package service

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
)

// SubjectInterface lets you work with common attributes from any kind of
// BastionZero subject.
type SubjectInterface interface {
	// GetID returns the subject's unique ID.
	GetID() string
	// GetSubjectType returns the subject's type.
	GetSubjectType() subjecttype.SubjectType
}
