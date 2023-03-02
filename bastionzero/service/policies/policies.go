package policies

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	policiesBasePath = "api/v2/policies"
)

// PoliciesService is an interface for interfacing with the policies endpoints
// of the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Policies
type PoliciesService interface {
	ListTargetConnectPolicies(ctx context.Context, opts *ListPolicyOptions) ([]TargetConnectPolicy, *http.Response, error)
	CreateTargetConnectPolicy(ctx context.Context, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error)
	GetTargetConnectPolicy(ctx context.Context, policyID string) (*TargetConnectPolicy, *http.Response, error)
	DeleteTargetConnectPolicy(ctx context.Context, policyID string) (*http.Response, error)
	ModifyTargetConnectPolicy(ctx context.Context, policyID string, policy *TargetConnectPolicy) (*TargetConnectPolicy, *http.Response, error)
}

// PoliciesServiceOp handles communication with the policies endpoints of the
// BastionZero API.
type PoliciesServiceOp client.Service

// Assert PoliciesServiceOp struct implements PoliciesService interface
var _ PoliciesService = &PoliciesServiceOp{}
