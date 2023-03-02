package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	targetConnectBasePath   = policiesBasePath + "/target-connect"
	targetConnectSinglePath = targetConnectBasePath + "/%s"
)

// TargetUser refers to the Unix user that a target connect policy applies to
type TargetUser struct {
	Username string `json:"userName"`
}

// Verb refers to an action allowed by a target connect policy
type Verb struct {
	Type VerbType `json:"type"`
}

// VerbType represents the type of target connect verb
type VerbType string

const (
	// Shell represents the ability to make a Shell connection
	Shell VerbType = "Shell"
	// FileTransfer represents the ability to upload/download files
	FileTransfer VerbType = "FileTransfer"
	// Tunnel represents the ability to make an SSH tunnel
	Tunnel VerbType = "Tunnel"
)

// TargetConnectPolicy represents a target connect policy. Target connect
// policies provide access to Bzero and DynamicAccessConfig targets.
type TargetConnectPolicy struct {
	// ID of the policy. Populated by the server
	ID string `json:"id,omitempty"`

	// User-initalized fields

	TimeExpires *service.Timestamp `json:"timeExpires,omitempty"`

	// User-mutable fields

	Name         string               `json:"name,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Subjects     *[]PolicySubject     `json:"subjects,omitempty"`
	Groups       *[]PolicyGroup       `json:"groups,omitempty"`
	Environments *[]PolicyEnvironment `json:"environments,omitempty"`
	Targets      *[]PolicyTarget      `json:"targets,omitempty"`
	TargetUsers  *[]TargetUser        `json:"targetUsers,omitempty"`
	Verbs        *[]Verb              `json:"verbs,omitempty"`
}

// ListTargetConnectPolicies lists all target connect policies
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/target-connect
func (s *PoliciesServiceOp) ListTargetConnectPolicies(ctx context.Context, opts *ListPolicyOptions) ([]TargetConnectPolicy, *http.Response, error) {
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

// CreateTargetConnectPolicy creates a new target connect policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/target-connect
func (s *PoliciesServiceOp) CreateTargetConnectPolicy(ctx context.Context, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error) {
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

// GetTargetConnectPolicy fetches the specified target connect policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/target-connect/-id-
func (s *PoliciesServiceOp) GetTargetConnectPolicy(ctx context.Context, policyID string) (*TargetConnectPolicy, *http.Response, error) {
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

// DeleteTargetConnectPolicy deletes the specified target connect policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/target-connect/-id-
func (s *PoliciesServiceOp) DeleteTargetConnectPolicy(ctx context.Context, policyID string) (*http.Response, error) {
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
func (s *PoliciesServiceOp) ModifyTargetConnectPolicy(ctx context.Context, policyID string, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error) {
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
