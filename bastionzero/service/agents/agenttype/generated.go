// Code generated by "string-enumer -t AgentType -o ./generated.go ."; DO NOT EDIT.
package agenttype

// validAgentTypeValues contains a map of all valid AgentType values for easy lookup
var validAgentTypeValues = map[AgentType]struct{}{
	Cluster: {},
	Linux:   {},
	Windows: {},
}

// Valid validates if a value is a valid AgentType
func (v AgentType) Valid() bool {
	_, ok := validAgentTypeValues[v]
	return ok
}

// AgentTypeValues returns a list of all (valid) AgentType values
func AgentTypeValues() []AgentType {
	return []AgentType{
		Cluster,
		Linux,
		Windows,
	}
}
