package policies

// ChildPolicy refers to another policy that a JIT policy applies to
type ChildPolicy struct {
	ID string `json:"id"`
	// PolicyType must be one of TargetConnect, Kubernetes or Proxy
	Type PolicyType `json:"type"`
	Name string     `json:"name"`
}
