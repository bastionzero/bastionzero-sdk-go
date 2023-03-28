package policytype

//go:generate go run github.com/lindell/string-enumer -t PolicyType -o ./generated.go .

// PolicyType represents the type of policy
type PolicyType string

const (
	// TargetConnect represents a target connect policy
	TargetConnect PolicyType = "TargetConnect"
	// OrganizationControls represents an organization controls policy
	OrganizationControls PolicyType = "OrganizationControls"
	// SessionRecording represents a session recording policy
	SessionRecording PolicyType = "SessionRecording"
	// Kubernetes represents a Kubernetes policy
	Kubernetes PolicyType = "Kubernetes"
	// Proxy represents a Proxy policy
	Proxy PolicyType = "Proxy"
	// JustInTime represents a JIT policy
	JustInTime PolicyType = "JustInTime"
)
