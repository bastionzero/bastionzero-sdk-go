package agenttype

//go:generate go run github.com/lindell/string-enumer -t AgentStatus -o ./generated.go .

// Where the agent is running. Either a kube cluster, a linux server, or a windows server
type AgentType string

const (
	// TODO: CWC-2744 deprecate Cluster in favor of Kubernetes
	Cluster AgentType = "Cluster"
	Linux   AgentType = "Linux"
	Windows AgentType = "Windows"
)
