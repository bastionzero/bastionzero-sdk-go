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

// PolicyTargetUser refers to the user that a policy applies to. When used in a
// target connect policy, it refers to a Unix username. When used in a proxy
// policy, it refers to a database user.
type PolicyTargetUser struct {
	Username string `json:"userName"`
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

// PolicyInterface lets you work with common policy attributes from any kind of
// BastionZero policy
type PolicyInterface interface {
	// GetID returns the policy's unique ID.
	GetID() string
	// GetTimeExpires returns the policy's expiration if set. Otherwise, returns
	// nil.
	GetTimeExpires() *service.Timestamp
	// GetName returns the the policy's name.
	GetName() string
	// GetDescription returns the policy's description if set. Otherwise,
	// returns an empty string.
	GetDescription() string
	// GetSubjects returns the policy's list of subjects that a policy applies
	// to if set. Otherwise, returns an empty slice.
	GetSubjects() []PolicySubject
	// GetGroups returns the policy's list of groups that a policy applies to if
	// set. Otherwise, returns an empty slice.
	GetGroups() []PolicyGroup
	// GetPolicyType returns the policy's type.
	GetPolicyType() PolicyType
}
