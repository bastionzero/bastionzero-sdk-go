package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/verbtype"
	"github.com/bastionzero/bastionzero-sdk-go/internal"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	targetConnectBasePath   = policiesBasePath + "/target-connect"
	targetConnectSinglePath = targetConnectBasePath + "/%s"
)

// Verb refers to an action allowed by a target connect policy
type Verb struct {
	Type verbtype.VerbType `json:"type"`
}

// TargetConnectPolicy represents a target connect policy. Target connect
// policies provide access to Bzero and DynamicAccessConfig targets.
type TargetConnectPolicy struct {
	*Policy

	Environments *[]Environment `json:"environments,omitempty"`
	Targets      *[]Target      `json:"targets,omitempty"`
	TargetUsers  *[]TargetUser  `json:"targetUsers,omitempty"`
	Verbs        *[]Verb        `json:"verbs,omitempty"`
}

// ListTargetConnectPolicies lists all target connect policies.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/target-connect
func (s *PoliciesService) ListTargetConnectPolicies(ctx context.Context, opts *ListPolicyOptions) ([]TargetConnectPolicy, *http.Response, error) {
	u := targetConnectBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]TargetConnectPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateTargetConnectPolicy creates a new target connect policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/target-connect
func (s *PoliciesService) CreateTargetConnectPolicy(ctx context.Context, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error) {
	u := targetConnectBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(TargetConnectPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetTargetConnectPolicy fetches the specified target connect policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/target-connect/-id-
func (s *PoliciesService) GetTargetConnectPolicy(ctx context.Context, policyID string) (*TargetConnectPolicy, *http.Response, error) {
	u := fmt.Sprintf(targetConnectSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(TargetConnectPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteTargetConnectPolicy deletes the specified target connect policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/target-connect/-id-
func (s *PoliciesService) DeleteTargetConnectPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(targetConnectSinglePath, policyID)
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

// ModifyTargetConnectPolicy updates a target connect policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/target-connect/-id-
func (s *PoliciesService) ModifyTargetConnectPolicy(ctx context.Context, policyID string, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error) {
	u := fmt.Sprintf(targetConnectSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(TargetConnectPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure TargetConnectPolicy implementation satisfies the expected interfaces.
var (
	// TargetConnectPolicy implements PolicyInterface
	_ PolicyInterface = &TargetConnectPolicy{}
)

func (p *TargetConnectPolicy) GetPolicyType() policytype.PolicyType { return policytype.TargetConnect }

func (p *TargetConnectPolicy) GetEnvironments() []Environment {
	if p.Environments == nil {
		return []Environment{}
	}
	return *p.Environments
}
func (p *TargetConnectPolicy) GetTargets() []Target {
	if p.Targets == nil {
		return []Target{}
	}
	return *p.Targets
}
func (p *TargetConnectPolicy) GetTargetUsers() []TargetUser {
	if p.TargetUsers == nil {
		return []TargetUser{}
	}
	return *p.TargetUsers
}
func (p *TargetConnectPolicy) GetVerbs() []Verb {
	if p.Verbs == nil {
		return []Verb{}
	}
	return *p.Verbs
}

func (p *TargetConnectPolicy) GetEnvironmentsAsStringList() []string {
	return internal.MapSlice(p.GetEnvironments(), func(e Environment) string { return e.ID })
}
func (p *TargetConnectPolicy) GetTargetUsersAsStringList() []string {
	return internal.MapSlice(p.GetTargetUsers(), func(e TargetUser) string { return e.Username })
}
func (p *TargetConnectPolicy) GetVerbsAsStringList() []string {
	return internal.MapSlice(p.GetVerbs(), func(e Verb) string { return string(e.Type) })
}
