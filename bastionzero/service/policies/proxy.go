package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	proxyBasePath   = policiesBasePath + "/proxy"
	proxySinglePath = proxyBasePath + "/%s"
)

// ProxyPolicy represents a proxy policy. Proxy policies provide access to DB
// and Web targets.
type ProxyPolicy struct {
	// ID of the policy. Populated by the server
	ID string `json:"id,omitempty"`

	// User-initialized fields

	TimeExpires *types.Timestamp `json:"timeExpires,omitempty"`

	// User-mutable fields

	Name         string         `json:"name,omitempty"`
	Description  *string        `json:"description,omitempty"`
	Subjects     *[]Subject     `json:"subjects,omitempty"`
	Groups       *[]Group       `json:"groups,omitempty"`
	Environments *[]Environment `json:"environments,omitempty"`
	Targets      *[]Target      `json:"targets,omitempty"`
	TargetUsers  *[]TargetUser  `json:"targetUsers,omitempty"`
}

// ListProxyPolicies lists all proxy policies.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/proxy
func (s *PoliciesService) ListProxyPolicies(ctx context.Context, opts *ListPolicyOptions) ([]ProxyPolicy, *http.Response, error) {
	u := proxyBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]ProxyPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateProxyPolicy creates a new proxy policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/proxy
func (s *PoliciesService) CreateProxyPolicy(ctx context.Context, policy *ProxyPolicy) (*ProxyPolicy, *http.Response, error) {
	u := proxyBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProxyPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetProxyPolicy fetches the specified proxy policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/proxy/-id-
func (s *PoliciesService) GetProxyPolicy(ctx context.Context, policyID string) (*ProxyPolicy, *http.Response, error) {
	u := fmt.Sprintf(proxySinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(ProxyPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteProxyPolicy deletes the specified proxy policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/proxy/-id-
func (s *PoliciesService) DeleteProxyPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(proxySinglePath, policyID)
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

// ModifyProxyPolicy updates a proxy policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/proxy/-id-
func (s *PoliciesService) ModifyProxyPolicy(ctx context.Context, policyID string, policy *ProxyPolicy) (*ProxyPolicy, *http.Response, error) {
	u := fmt.Sprintf(proxySinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProxyPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure ProxyPolicy implementation satisfies the expected interfaces.
var (
	// ProxyPolicy implements PolicyInterface
	_ PolicyInterface = &ProxyPolicy{}
)

func (p *ProxyPolicy) GetID() string                    { return p.ID }
func (p *ProxyPolicy) GetTimeExpires() *types.Timestamp { return p.TimeExpires }
func (p *ProxyPolicy) GetName() string                  { return p.Name }
func (p *ProxyPolicy) GetDescription() string {
	if p.Description == nil {
		return ""
	}
	return *p.Description
}
func (p *ProxyPolicy) GetSubjects() []Subject {
	if p.Subjects == nil {
		return []Subject{}
	}
	return *p.Subjects
}
func (p *ProxyPolicy) GetGroups() []Group {
	if p.Groups == nil {
		return []Group{}
	}
	return *p.Groups
}
func (p *ProxyPolicy) GetPolicyType() policytype.PolicyType { return policytype.Proxy }

func (p *ProxyPolicy) GetEnvironments() []Environment {
	if p.Environments == nil {
		return []Environment{}
	}
	return *p.Environments
}
func (p *ProxyPolicy) GetTargets() []Target {
	if p.Targets == nil {
		return []Target{}
	}
	return *p.Targets
}
func (p *ProxyPolicy) GetTargetUsers() []TargetUser {
	if p.TargetUsers == nil {
		return []TargetUser{}
	}
	return *p.TargetUsers
}

func (p *ProxyPolicy) GetEnvironmentsAsStringList() []string {
	return internal.MapSlice(p.GetEnvironments(), func(e Environment) string { return e.ID })
}
func (p *ProxyPolicy) GetTargetUsersAsStringList() []string {
	return internal.MapSlice(p.GetTargetUsers(), func(e TargetUser) string { return e.Username })
}
