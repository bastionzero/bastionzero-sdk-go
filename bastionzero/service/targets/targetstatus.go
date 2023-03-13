package targets

//go:generate go run github.com/dmarkham/enumer -type=TargetStatus -json

// TargetStatus represents the status of a target
type TargetStatus int

const (
	// NotActivated represents a target that has not been activated
	NotActivated TargetStatus = iota
	// Offline represents a target that is offline
	Offline
	// Online represents a target that is online
	Online
	// Terminated represents a target that has been deleted
	Terminated
	// Error represents a target that has entered an errored state
	Error
	// Restarting represents a target that is currently restarting due to
	// receiving a restart request issued via the BastionZero API
	Restarting
)
