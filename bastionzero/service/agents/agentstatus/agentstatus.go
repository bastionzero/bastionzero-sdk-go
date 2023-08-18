package agentstatus

//go:generate go run github.com/lindell/string-enumer -t AgentStatus -o ./generated.go .

// AgentStatus represents the status of an agent
type AgentStatus string

const (
	// NotActivated represents an agent that has not been activated
	NotActivated AgentStatus = "NotActivated"
	// Offline represents an agent that is offline
	Offline AgentStatus = "Offline"
	// Online represents an agent that is online
	Online AgentStatus = "Online"
	// Terminated represents an agent that has been deleted
	Terminated AgentStatus = "Terminated"
	// Error represents an agent that has entered an errored state
	Error AgentStatus = "Error"
	// Restarting represents an agent that is currently restarting due to
	// receiving a restart request issued via the BastionZero API
	Restarting AgentStatus = "Restarting"
)
