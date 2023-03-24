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

// ServiceAccount is a Google, Azure, or generic service account that integrates
// with BastionZero by sharing its JSON Web Key Set (JWKS) URL. The headless
// authentication closely follows the OpenID Connect (OIDC) protocol. The JWKS
// contains the public key from a public/private key pair that you must
// generate. You use the private key to sign the service account’s identity and
// access tokens, and then BastionZero uses the public key within the JWKS URL
// to validate the service account.
type ServiceAccount struct {
	*service.Subject

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

// Ensure ServiceAccount implementation satisfies the expected interfaces.
var (
	// ServiceAccount implements SubjectInterface
	_ service.SubjectInterface = &ServiceAccount{}
)

func (s *ServiceAccount) GetSubjectType() subjecttype.SubjectType { return subjecttype.ServiceAccount }
