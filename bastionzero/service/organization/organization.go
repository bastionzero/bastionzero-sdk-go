package organization

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	organizationBasePath    = "api/v2/organization"
	groupsBasePath          = organizationBasePath + "/groups"
	registrationKeyBasePath = organizationBasePath + "/registration-key"
)

// OrganizationService handles communication with the organization endpoints of
// the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Organization
type OrganizationService client.Service

// Organization represents a BastionZero organization.
type Organization struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	IsSingleUserOrganization bool            `json:"isSingleUserOrganization"`
	TimeCreated              types.Timestamp `json:"timeCreated"`
}

// BZCertValidationInfo contains information needed to validate BZCerts issued
// for one's BastionZero organization.
type BZCertValidationInfo struct {
	OrgIdpProvider string `json:"orgIdpProvider"`
	OrgIdpIssuerId string `json:"orgIdpIssuerId"`
}

// Group is an Identity provider (IdP) group synced to BastionZero. A group
// contains users from an organization.
type Group struct {
	ID   string `json:"idPGroupId"`
	Name string `json:"name"`
}

// SlackIntegration contains details about a Slack integration with one's
// BastionZero organization.
type SlackIntegration struct {
	TeamName       string          `json:"teamName"`
	AdminEmail     string          `json:"adminEmail"`
	CreationDate   types.Timestamp `json:"creationDate"`
	LastUpdateDate types.Timestamp `json:"lastUpdateDate"`
}

// RegistrationKeySettings contains information about one's BastionZero
// organization's registration key settings used in the autodiscovery process.
type RegistrationKeySettings struct {
	GlobalRegistrationKeyEnforced bool    `json:"globalRegistrationKeyEnforced"`
	DefaultGlobalRegistrationKey  *string `json:"defaultGlobalRegistrationKey"`
}

// EnableGlobalRegistrationKeyRequest is used to enable global registration key
// enforcement
type EnableGlobalRegistrationKeyRequest struct {
	// DefaultRegistrationKeyId is the identifier for the registration key that
	// should be set as the default global registration key.
	DefaultRegistrationKeyId string `json:"defaultRegistrationKeyId"`
}

// IdentityProvider represents the identity provider (IdP) for one's BastionZero
// organization.
type IdentityProvider struct {
	IdentityProviderType string `json:"identityProviderType"`
	IdentityProviderId   string `json:"identityProviderId"`
}

// GetOrganization gets information about your BastionZero organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization
func (s *OrganizationService) GetOrganization(ctx context.Context) (*Organization, *http.Response, error) {
	u := organizationBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.Client.Do(req, org)
	if err != nil {
		return nil, resp, err
	}

	return org, resp, nil
}

// GetBZCertValidationInfo gets information needed to validate BZCerts issued
// for your BastionZero organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/bzcert-validation-info
func (s *OrganizationService) GetBZCertValidationInfo(ctx context.Context) (*BZCertValidationInfo, *http.Response, error) {
	u := organizationBasePath + "/bzcert-validation-info"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	validationInfo := new(BZCertValidationInfo)
	resp, err := s.Client.Do(req, validationInfo)
	if err != nil {
		return nil, resp, err
	}

	return validationInfo, resp, nil
}

// ListGroups lists the groups configured (synced from IdP) for your organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/groups
func (s *OrganizationService) ListGroups(ctx context.Context) ([]Group, *http.Response, error) {
	u := groupsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupList := new([]Group)
	resp, err := s.Client.Do(req, groupList)
	if err != nil {
		return nil, resp, err
	}

	return *groupList, resp, nil
}

// FetchGroups fetches the groups for your organization by querying the
// configured identity providers directly.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/organization/groups/fetch
func (s *OrganizationService) FetchGroups(ctx context.Context) ([]Group, *http.Response, error) {
	u := groupsBasePath + "/fetch"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupList := new([]Group)
	resp, err := s.Client.Do(req, groupList)
	if err != nil {
		return nil, resp, err
	}

	return *groupList, resp, nil
}

// FetchUserGroups fetches the groups for the specified user by querying the
// configured identity providers directly.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/organization/groups-memberships/fetch/-id-
func (s *OrganizationService) FetchUserGroups(ctx context.Context, userID string) ([]Group, *http.Response, error) {
	u := fmt.Sprintf(organizationBasePath+"/groups-memberships/fetch/%s", userID)
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	groupList := new([]Group)
	resp, err := s.Client.Do(req, groupList)
	if err != nil {
		return nil, resp, err
	}

	return *groupList, resp, nil
}

// DeleteIdpGroupCredentials deletes the credentials metadata used to fetch groups from the Identity Provider.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/organization/groups/credentials
func (s *OrganizationService) DeleteIdpGroupCredentials(ctx context.Context) (*http.Response, error) {
	u := groupsBasePath + "/credentials"
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

// InvalidateKeycloakProviderCache invalidates the Keycloak provider cache entry.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/organization/invalidate-keycloak
func (s *OrganizationService) InvalidateKeycloakProviderCache(ctx context.Context) (*http.Response, error) {
	u := organizationBasePath + "/invalidate-keycloak"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetSlackIntegration gets information about the Slack integration with your BastionZero organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/integrations/slack
func (s *OrganizationService) GetSlackIntegration(ctx context.Context) (*SlackIntegration, *http.Response, error) {
	u := organizationBasePath + "/integrations/slack"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	integration := new(SlackIntegration)
	resp, err := s.Client.Do(req, integration)
	if err != nil {
		return nil, resp, err
	}

	return integration, resp, nil
}

// GetRegistrationKeySettings gets information about the registration key
// settings for your BastionZero organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/registration-key/settings
func (s *OrganizationService) GetRegistrationKeySettings(ctx context.Context) (*RegistrationKeySettings, *http.Response, error) {
	u := registrationKeyBasePath + "/settings"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	settings := new(RegistrationKeySettings)
	resp, err := s.Client.Do(req, settings)
	if err != nil {
		return nil, resp, err
	}

	return settings, resp, nil
}

// EnableGlobalRegistrationKey enables global registration key enforcement for
// your BastionZero organization. If enabled, the global registration key is the
// only key that can be used for registering targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/organization/registration-key/enable-enforce-global-key
func (s *OrganizationService) EnableGlobalRegistrationKey(ctx context.Context, request *EnableGlobalRegistrationKeyRequest) (*RegistrationKeySettings, *http.Response, error) {
	u := registrationKeyBasePath + "/enable-enforce-global-key"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	settings := new(RegistrationKeySettings)
	resp, err := s.Client.Do(req, settings)
	if err != nil {
		return nil, resp, err
	}

	return settings, resp, nil
}

// DisableGlobalRegistrationKey disables global registration key enforcement for
// your BastionZero organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/organization/registration-key/disable-enforce-global-key
func (s *OrganizationService) DisableGlobalRegistrationKey(ctx context.Context) (*RegistrationKeySettings, *http.Response, error) {
	u := registrationKeyBasePath + "/disable-enforce-global-key"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	settings := new(RegistrationKeySettings)
	resp, err := s.Client.Do(req, settings)
	if err != nil {
		return nil, resp, err
	}

	return settings, resp, nil
}

// GetIdentityProvider gets identity provider details for your BastionZero
// organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/organization/identity-provider
func (s *OrganizationService) GetIdentityProvider(ctx context.Context) (*IdentityProvider, *http.Response, error) {
	u := organizationBasePath + "/identity-provider"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	idpDetails := new(IdentityProvider)
	resp, err := s.Client.Do(req, idpDetails)
	if err != nil {
		return nil, resp, err
	}

	return idpDetails, resp, nil
}
