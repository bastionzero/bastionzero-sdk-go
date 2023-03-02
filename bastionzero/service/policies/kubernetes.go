package policies

// ClusterUser refers to the Kubernetes subject user that a Kubernetes policy
// applies to
type ClusterUser struct {
	Name string `json:"name"`
}

// ClusterGroup refers to the Kubernetes subject group that a Kubernetes policy
// applies to
type ClusterGroup struct {
	Name string `json:"name"`
}

// Cluster refers to the BastionZero Cluster target that a Kubernetes policy
// applies to
type Cluster struct {
	ID string `json:"id"`
}
