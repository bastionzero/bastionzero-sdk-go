package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	orgControlsBasePath   = policiesBasePath + "/organization-controls"
	orgControlsSinglePath = orgControlsBasePath + "/%s"
)

// OrganizationControlsPolicy represents an OrganizationControls policy. This policy controls global mfa.
type OrganizationControlsPolicy struct {
	Policy

	MFAEnabled  *bool `json:"mfaEnabled,omitempty"`
	MFADuration *int  `json:"mfaDuration,omitempty"`
}

// ListOrganizationControlsPolicies lists all Organization Controls policies
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/organization-controls
func (s *PoliciesService) ListOrganizationControlsPolicies(ctx context.Context, opts *ListPolicyOptions) ([]OrganizationControlsPolicy, *http.Response, error) {
	u := orgControlsBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]OrganizationControlsPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateOrganizationControlsPolicy creates a new OrganizationControls policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/organization-controls
func (s *PoliciesService) CreateOrganizationControlsPolicy(ctx context.Context, policy *OrganizationControlsPolicy) (*OrganizationControlsPolicy, *http.Response, error) {
	u := orgControlsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(OrganizationControlsPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetOrganizationControlsPolicyPolicy fetches the specified OrganizationControlsPolicy policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/organization-controls/-id-
func (s *PoliciesService) GetOrganizationControlsPolicyPolicy(ctx context.Context, policyID string) (*OrganizationControlsPolicy, *http.Response, error) {
	u := fmt.Sprintf(orgControlsSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(OrganizationControlsPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteOrganizationControlsPolicy deletes the specified OrganizationControls policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/organization-controls/-id-
func (s *PoliciesService) DeleteOrganizationControlsPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(orgControlsSinglePath, policyID)
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

// ModifyOrganizationControlsPolicy updates a OrganizationControls policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/organization-controls/-id-
func (s *PoliciesService) ModifyOrganizationControlsPolicy(ctx context.Context, policyID string, policy *OrganizationControlsPolicy) (*OrganizationControlsPolicy, *http.Response, error) {
	u := fmt.Sprintf(orgControlsSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(OrganizationControlsPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure OrganizationControlsPolicy implementation satisfies the expected interfaces.
var (
	// OrganizationControlsPolicy implements PolicyInterface
	_ PolicyInterface = &OrganizationControlsPolicy{}
)

func (p *OrganizationControlsPolicy) GetPolicyType() policytype.PolicyType {
	return policytype.OrganizationControls
}

func (p *OrganizationControlsPolicy) GetMFAEnabled() bool {
	if p.MFAEnabled == nil {
		return false
	}
	return *p.MFAEnabled
}

func (p *OrganizationControlsPolicy) GetMFADuration() int {
	if p.MFADuration == nil {
		return 0
	}
	return *p.MFADuration
}
