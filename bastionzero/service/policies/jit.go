package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	jitBasePath   = policiesBasePath + "/just-in-time"
	jitSinglePath = jitBasePath + "/%s"
)

// CreateJITPolicyRequest is used to create a new JIT policy
type CreateJITPolicyRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	TimeExpires *types.Timestamp `json:"timeExpires,omitempty"`
	Subjects    []Subject        `json:"subjects"`
	Groups      []Group          `json:"groups"`
	// ChildPolicies is a list of policy IDs that this JIT policy applies to
	ChildPolicies []string `json:"childPolicies"`
	// AutomaticallyApproved determines whether the creation of the policies
	// will be automatically approved or based on request and approval from
	// reviewers.
	AutomaticallyApproved bool `json:"automaticallyApproved"`
	// Duration is the amount of time (in minutes) after which the access
	// granted by this JIT policy will expire.
	Duration uint `json:"duration"`
}

// ModifyJITPolicyRequest is used to modify a JIT policy
type ModifyJITPolicyRequest struct {
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Subjects    *[]Subject `json:"subjects,omitempty"`
	Groups      *[]Group   `json:"groups,omitempty"`
	// ChildPolicies is a list of policy IDs that this JIT policy applies to
	ChildPolicies *[]string `json:"childPolicies,omitempty"`
	// AutomaticallyApproved determines whether the creation of the policies
	// will be automatically approved or based on request and approval from
	// reviewers.
	AutomaticallyApproved *bool `json:"automaticallyApproved,omitempty"`
	// Duration is the amount of time (in minutes) after which the access
	// granted by this JIT policy will expire.
	Duration *uint `json:"duration,omitempty"`
}

// ChildPolicy refers to another policy that a JIT policy applies to
type ChildPolicy struct {
	ID string `json:"id"`
	// Type is one of TargetConnect, Kubernetes or Proxy
	Type policytype.PolicyType `json:"type"`
	Name string                `json:"name"`
}

// JITPolicy represents a just in time policy. Just in time policies provide
// just in time access to targets.
type JITPolicy struct {
	ID          string           `json:"id"`
	TimeExpires *types.Timestamp `json:"timeExpires"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Subjects    []Subject        `json:"subjects"`
	Groups      []Group          `json:"groups"`
	// ChildPolicies is a list of policies that this JIT policy applies to
	ChildPolicies []ChildPolicy `json:"childPolicies"`
	// AutomaticallyApproved determines whether the creation of the policies
	// will be automatically approved or based on request and approval from
	// reviewers.
	AutomaticallyApproved bool `json:"automaticallyApproved"`
	// Duration is the amount of time (in minutes) after which the access
	// granted by this JIT policy will expire.
	Duration uint `json:"duration"`
}

// ListJITPolicies lists all JIT policies.
//
// BastionZero API docs:
// https://cloud.bastionzero.com/api/#get-/api/v2/policies/just-in-time
func (s *PoliciesService) ListJITPolicies(ctx context.Context, opts *ListPolicyOptions) ([]JITPolicy, *http.Response, error) {
	u := jitBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]JITPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateJITPolicy creates a new JIT policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/just-in-time
func (s *PoliciesService) CreateJITPolicy(ctx context.Context, request *CreateJITPolicyRequest) (*JITPolicy, *http.Response, error) {
	u := jitBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	p := new(JITPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetJITPolicy fetches the specified JIT policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/just-in-time/-id-
func (s *PoliciesService) GetJITPolicy(ctx context.Context, policyID string) (*JITPolicy, *http.Response, error) {
	u := fmt.Sprintf(jitSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(JITPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteJITPolicy deletes the specified JIT policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/just-in-time/-id-
func (s *PoliciesService) DeleteJITPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(jitSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ModifyJITPolicy updates a JIT policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/just-in-time/-id-
func (s *PoliciesService) ModifyJITPolicy(ctx context.Context, policyID string, request *ModifyJITPolicyRequest) (*JITPolicy, *http.Response, error) {
	u := fmt.Sprintf(jitSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	p := new(JITPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure JITPolicy implementation satisfies the expected interfaces.
var (
	// JITPolicy implements PolicyInterface
	_ PolicyInterface = &JITPolicy{}
)

func (p *JITPolicy) GetID() string                    { return p.ID }
func (p *JITPolicy) GetTimeExpires() *types.Timestamp { return p.TimeExpires }
func (p *JITPolicy) GetName() string                  { return p.Name }
func (p *JITPolicy) GetDescription() string {
	return p.Description
}
func (p *JITPolicy) GetSubjects() []Subject {
	return p.Subjects
}
func (p *JITPolicy) GetGroups() []Group {
	return p.Groups
}
func (p *JITPolicy) GetPolicyType() policytype.PolicyType { return policytype.JustInTime }

func (p *JITPolicy) GetChildPolicies() []ChildPolicy {
	return p.ChildPolicies
}
func (p *JITPolicy) GetAutomaticallyApproved() bool {
	return p.AutomaticallyApproved
}
func (p *JITPolicy) GetDuration() uint {
	return p.Duration
}
