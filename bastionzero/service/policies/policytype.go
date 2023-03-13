package policies

//go:generate go run github.com/dmarkham/enumer -type=PolicyType -json

// PolicyType represents the type of policy
type PolicyType int

const (
	// TargetConnect represents a target connect policy
	TargetConnect PolicyType = iota
	// OrganizationControls represents an organization controls policy
	OrganizationControls
	// SessionRecording represents a session recording policy
	SessionRecording
	// Kubernetes represents a Kubernetes policy
	Kubernetes
	// Proxy represents a Proxy policy
	Proxy
	// JustInTime represents a JIT policy
	JustInTime
)
