package roletype

//go:generate go run github.com/lindell/string-enumer -t RoleType -o ./generated.go .

// RoleType represents the type of role a subject is assigned
type RoleType string

const (
	// User denotes the user role
	User RoleType = "User"
	// Admin denotes the admin role
	Admin RoleType = "Admin"
)
