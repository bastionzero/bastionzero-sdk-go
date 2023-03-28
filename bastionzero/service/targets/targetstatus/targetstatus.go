package targetstatus

//go:generate go run github.com/lindell/string-enumer -t TargetStatus -o ./generated.go .

// TargetStatus represents the status of a target
type TargetStatus string

const (
	// NotActivated represents a target that has not been activated
	NotActivated TargetStatus = "NotActivated"
	// Offline represents a target that is offline
	Offline TargetStatus = "Offline"
	// Online represents a target that is online
	Online TargetStatus = "Online"
	// Terminated represents a target that has been deleted
	Terminated TargetStatus = "Terminated"
	// Error represents a target that has entered an errored state
	Error TargetStatus = "Error"
	// Restarting represents a target that is currently restarting due to
	// receiving a restart request issued via the BastionZero API
	Restarting TargetStatus = "Restarting"
)
