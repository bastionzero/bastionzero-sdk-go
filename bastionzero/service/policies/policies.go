package policies

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
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

// Subject refers to the subject that a policy applies to
type Subject struct {
	ID   string                  `json:"id"`
	Type subjecttype.SubjectType `json:"type"`
}

// Group refers to the IdP group that a policy applies to
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Target refers to the BastionZero target that a policy applies to
type Target struct {
	ID   string                `json:"id"`
	Type targettype.TargetType `json:"type"`
}

// Environment refers to the BastionZero environment that a policy applies
// to
type Environment struct {
	ID string `json:"id"`
}

// TargetUser refers to the user that a policy applies to. When used in a
// target connect policy, it refers to a Unix username. When used in a proxy
// policy, it refers to a database user.
type TargetUser struct {
	Username string `json:"userName"`
}

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
	GetTimeExpires() *types.Timestamp
	// GetName returns the the policy's name.
	GetName() string
	// GetDescription returns the policy's description if set. Otherwise,
	// returns an empty string.
	GetDescription() string
	// GetSubjects returns the policy's list of subjects that the policy applies
	// to if set. Otherwise, returns an empty slice.
	GetSubjects() []Subject
	// GetGroups returns the policy's list of groups that the policy applies to
	// if set. Otherwise, returns an empty slice.
	GetGroups() []Group
	// GetPolicyType returns the policy's type.
	GetPolicyType() policytype.PolicyType
}
