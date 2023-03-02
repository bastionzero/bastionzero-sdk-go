package policies

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	policiesBasePath = "api/v2/policies"
)

// PoliciesService handles communication with the policies endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Policies
type PoliciesService client.Service

// PolicySubject refers to the subject that a policy applies to
type PolicySubject struct {
	ID   string              `json:"id"`
	Type service.SubjectType `json:"type"`
}

// PolicyGroup refers to the IdP group that a policy applies to
type PolicyGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PolicyTarget refers to the BastionZero target that a policy applies to
type PolicyTarget struct {
	ID   string             `json:"id"`
	Type service.TargetType `json:"type"`
}

// PolicyEnvironment refers to the BastionZero environment that a policy applies
// to
type PolicyEnvironment struct {
	ID string `json:"id"`
}

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

// ListPolicyOptions specifies the optional parameters to various List methods
// involving policy
type ListPolicyOptions struct {
	// Subjects is a comma-separated list of subject IDs. Filters the list of
	// policies to only those that contain the provided subject(s).
	Subjects string `url:"subjects,omitempty"`

	// Groups is a comma-separated list of group IDs. Filters the list of
	// policies to only those that contain the provided group(s).
	Groups string `url:"groups,omitempty"`
}
