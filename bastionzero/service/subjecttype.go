package service

//go:generate go run github.com/dmarkham/enumer -type=SubjectType -json

// SubjectType represents the type of user interacting with the BastionZero API.
type SubjectType int

const (
	// User represents a user belonging to an IdP's organization
	User SubjectType = iota
	// ApiKey represents an API key created via the BastionZero API
	ApiKey
	// ServiceAccount represents a service account identity managed by some IdP
	ServiceAccount
)
