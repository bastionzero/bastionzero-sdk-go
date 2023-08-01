// Code generated by "string-enumer -t AgentStatus -o ./generated.go ."; DO NOT EDIT.
package agentstatus

// validAgentStatusValues contains a map of all valid AgentStatus values for easy lookup
var validAgentStatusValues = map[AgentStatus]struct{}{
	NotActivated: {},
	Offline:      {},
	Online:       {},
	Terminated:   {},
	Error:        {},
	Restarting:   {},
}

// Valid validates if a value is a valid AgentStatus
func (v AgentStatus) Valid() bool {
	_, ok := validAgentStatusValues[v]
	return ok
}

// AgentStatusValues returns a list of all (valid) AgentStatus values
func AgentStatusValues() []AgentStatus {
	return []AgentStatus{
		NotActivated,
		Offline,
		Online,
		Terminated,
		Error,
		Restarting,
	}
}
