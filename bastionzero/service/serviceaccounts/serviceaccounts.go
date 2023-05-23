package serviceaccounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	serviceaccountsBasePath   = "api/v2/service-accounts"
	serviceaccountsSinglePath = serviceaccountsBasePath + "/%s"
)

// ServiceAccountsService handles communication with the service accounts endpoints of
// the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Service-Accounts
type ServiceAccountsService client.Service

// CreateServiceAccountRequest is used to create a new service account
type CreateServiceAccountRequest struct {
	Email          string `json:"email"`
	JwksURL        string `json:"jwksUrl"`
	JwksURLPattern string `json:"jwksUrlPattern"`
	ExternalId     string `json:"externalId"`
}

// CreateServiceAccountResponse is the response returned if a service account is
// successfully created
type CreateServiceAccountResponse struct {
	ServiceAccountSummary ServiceAccount `json:"serviceAccountSummary"`
	MFASecret             string         `json:"mfaSecret"`
}

// ModifyServiceAccountRequest is used to modify a service account
type ModifyServiceAccountRequest struct {
	IsAdmin *bool `json:"isAdmin,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// ServiceAccount is a Google, Azure, or generic service account that integrates
// with BastionZero by sharing its JSON Web Key Set (JWKS) URL. The headless
// authentication closely follows the OpenID Connect (OIDC) protocol. The JWKS
// contains the public key from a public/private key pair that you must
// generate. You use the private key to sign the service account’s identity and
// access tokens, and then BastionZero uses the public key within the JWKS URL
// to validate the service account.
type ServiceAccount struct {
	service.Subject

	OrganizationID string           `json:"organizationId"`
	Email          string           `json:"email"`
	ExternalID     string           `json:"externalId"`
	JwksURL        string           `json:"jwksUrl"`
	JwksURLPattern string           `json:"jwksUrlPattern"`
	IsAdmin        bool             `json:"isAdmin"`
	TimeCreated    types.Timestamp  `json:"timeCreated"`
	LastLogin      *types.Timestamp `json:"lastLogin"`
	CreatedBy      string           `json:"createdBy"`
	Enabled        bool             `json:"enabled"`
}

// ListServiceAccounts lists all service accounts for your organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/service-accounts
func (s *ServiceAccountsService) ListServiceAccounts(ctx context.Context) ([]ServiceAccount, *http.Response, error) {
	u := serviceaccountsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	serviceAccountList := new([]ServiceAccount)
	resp, err := s.Client.Do(req, serviceAccountList)
	if err != nil {
		return nil, resp, err
	}

	return *serviceAccountList, resp, nil
}

// CreateServiceAccount creates a new service account.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/service-accounts
func (s *ServiceAccountsService) CreateServiceAccount(ctx context.Context, request *CreateServiceAccountRequest) (*CreateServiceAccountResponse, *http.Response, error) {
	u := serviceaccountsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createServiceAccountResponse := new(CreateServiceAccountResponse)
	resp, err := s.Client.Do(req, createServiceAccountResponse)
	if err != nil {
		return nil, resp, err
	}

	return createServiceAccountResponse, resp, nil
}

// GetServiceAccount fetches the specified service account by ID.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/service-accounts/-id-
func (s *ServiceAccountsService) GetServiceAccount(ctx context.Context, id string) (*ServiceAccount, *http.Response, error) {
	u := fmt.Sprintf(serviceaccountsSinglePath, id)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	serviceAccount := new(ServiceAccount)
	resp, err := s.Client.Do(req, serviceAccount)
	if err != nil {
		return nil, resp, err
	}

	return serviceAccount, resp, nil
}

// ModifyServiceAccount updates a service account.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/service-accounts/-id-
func (s *ServiceAccountsService) ModifyServiceAccount(ctx context.Context, serviceAccountID string, request *ModifyServiceAccountRequest) (*ServiceAccount, *http.Response, error) {
	u := fmt.Sprintf(serviceaccountsSinglePath, serviceAccountID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	serviceAccount := new(ServiceAccount)
	resp, err := s.Client.Do(req, serviceAccount)
	if err != nil {
		return nil, resp, err
	}

	return serviceAccount, resp, nil
}

// Me fetches your service account information (current subject).
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/service-accounts/me
func (s *ServiceAccountsService) Me(ctx context.Context) (*ServiceAccount, *http.Response, error) {
	u := serviceaccountsBasePath + "/me"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	serviceAccount := new(ServiceAccount)
	resp, err := s.Client.Do(req, serviceAccount)
	if err != nil {
		return nil, resp, err
	}

	return serviceAccount, resp, nil
}

// InvalidateJwksURLCache invalidates the Jwks URL cache of the specified
// service account.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/service-accounts/invalidate-cache/-id-
func (s *ServiceAccountsService) InvalidateJwksURLCache(ctx context.Context, serviceAccountID string) (*http.Response, error) {
	u := fmt.Sprintf(serviceaccountsBasePath+"/invalidate-cache/%s", serviceAccountID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Ensure ServiceAccount implementation satisfies the expected interfaces.
var (
	// ServiceAccount implements SubjectInterface
	_ service.SubjectInterface = &ServiceAccount{}
)

func (s *ServiceAccount) GetSubjectType() subjecttype.SubjectType { return subjecttype.ServiceAccount }
