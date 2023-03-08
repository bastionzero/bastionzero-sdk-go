package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	jitBasePath   = policiesBasePath + "/just-in-time"
	jitSinglePath = jitBasePath + "/%s"
)

// ChildPolicy refers to another policy that a JIT policy applies to
type ChildPolicy struct {
	ID string `json:"id"`
	// Type must be one of TargetConnect, Kubernetes or Proxy
	Type PolicyType `json:"type"`
	Name string     `json:"name"`
}

// JITPolicy represents a just in time policy. Just in time policies provide
// just in time access to targets.
type JITPolicy struct {
	// ID of the policy. Populated by the server
	ID string `json:"id,omitempty"`

	// User-initialized fields

	TimeExpires *service.Timestamp `json:"timeExpires,omitempty"`

	// User-mutable fields

	Name          string           `json:"name,omitempty"`
	Description   *string          `json:"description,omitempty"`
	Subjects      *[]PolicySubject `json:"subjects,omitempty"`
	Groups        *[]PolicyGroup   `json:"groups,omitempty"`
	ChildPolicies *[]ChildPolicy   `json:"childPolicies,omitempty"`
	// AutomaticallyApproved determines whether the creation of the policies
	// will be automatically approved or based on request and approval from
	// reviewers.
	AutomaticallyApproved *bool `json:"automaticallyApproved,omitempty"`
	// Duration is the amount of time (in minutes) after which the access
	// granted by this JIT policy will expire.
	Duration *uint `json:"duration,omitempty"`
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
func (s *PoliciesService) CreateJITPolicy(ctx context.Context, policy *JITPolicy) (*JITPolicy, *http.Response, error) {
	u := jitBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
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

// ModifyJITPolicy updates a JIT policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/just-in-time/-id-
func (s *PoliciesService) ModifyJITPolicy(ctx context.Context, policyID string, policy *JITPolicy) (*JITPolicy, *http.Response, error) {
	u := fmt.Sprintf(jitSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
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

func (p *JITPolicy) GetID() string                      { return p.ID }
func (p *JITPolicy) GetTimeExpires() *service.Timestamp { return p.TimeExpires }
func (p *JITPolicy) GetName() string                    { return p.Name }
func (p *JITPolicy) GetDescription() string {
	if p.Description == nil {
		return ""
	}
	return *p.Description
}
func (p *JITPolicy) GetSubjects() []PolicySubject {
	if p.Subjects == nil {
		return []PolicySubject{}
	}
	return *p.Subjects
}
func (p *JITPolicy) GetGroups() []PolicyGroup {
	if p.Groups == nil {
		return []PolicyGroup{}
	}
	return *p.Groups
}
func (p *JITPolicy) GetPolicyType() PolicyType { return JustInTime }
