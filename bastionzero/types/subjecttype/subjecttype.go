package subjecttype

//go:generate go run github.com/lindell/string-enumer -t SubjectType -o ./generated.go .

// TODO-Yuval: Move this entire package/folder to service/subjects batched in a
// breaking change release

// SubjectType represents the type of user interacting with the BastionZero API.
type SubjectType string

const (
	// User represents a user belonging to an IdP's organization
	User SubjectType = "User"
	// ApiKey represents an API key created via the BastionZero API
	ApiKey SubjectType = "ApiKey"
	// ServiceAccount represents a service account identity managed by some IdP
	ServiceAccount SubjectType = "ServiceAccount"
)
