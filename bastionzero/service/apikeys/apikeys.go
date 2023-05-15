package apikeys

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
	apikeysBasePath   = "api/v2/api-keys"
	apikeysSinglePath = apikeysBasePath + "/%s"
)

// ApiKeysService handles communication with the api keys endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--API-Keys
type ApiKeysService client.Service

// CreateGlobalApiKeyRequest is used to create a new globally scoped API key
type CreateGlobalApiKeyRequest struct {
	Name              string `json:"name,omitempty"`
	IsRegistrationKey bool   `json:"isRegistrationKey"`
}

// CreateGlobalApiKeyResponse is the response returned if a globally scoped API
// key is successfully created
type CreateGlobalApiKeyResponse struct {
	ApiKeyDetails ApiKey `json:"apiKeyDetails"`
	Secret        string `json:"secret"`
}

// ModifyApiKeyRequest is used to modify an API key
type ModifyApiKeyRequest struct {
	Name *string `json:"name,omitempty"`
}

// ApiKey is an API key that provides programmatic access to select BastionZero
// API endpoints
type ApiKey struct {
	service.Subject

	Name              string          `json:"name"`
	TimeCreated       types.Timestamp `json:"timeCreated"`
	IsRegistrationKey bool            `json:"isRegistrationKey"`
}

// ListGlobalApiKeys lists all globally scoped API keys.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/api-keys
func (s *ApiKeysService) ListGlobalApiKeys(ctx context.Context) ([]ApiKey, *http.Response, error) {
	u := apikeysBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	globalApiKeyList := new([]ApiKey)
	resp, err := s.Client.Do(req, globalApiKeyList)
	if err != nil {
		return nil, resp, err
	}

	return *globalApiKeyList, resp, nil
}

// CreateGlobalApiKey creates a new globally scoped API key.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/api-keys
func (s *ApiKeysService) CreateGlobalApiKey(ctx context.Context, request *CreateGlobalApiKeyRequest) (*CreateGlobalApiKeyResponse, *http.Response, error) {
	u := apikeysBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	createGlobalApiKeyResponse := new(CreateGlobalApiKeyResponse)
	resp, err := s.Client.Do(req, createGlobalApiKeyResponse)
	if err != nil {
		return nil, resp, err
	}

	return createGlobalApiKeyResponse, resp, nil
}

// GetApiKey fetches the specified API key.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/api-keys/-id-
func (s *ApiKeysService) GetApiKey(ctx context.Context, apiKeyID string) (*ApiKey, *http.Response, error) {
	u := fmt.Sprintf(apikeysSinglePath, apiKeyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	apiKey := new(ApiKey)
	resp, err := s.Client.Do(req, apiKey)
	if err != nil {
		return nil, resp, err
	}

	return apiKey, resp, nil
}

// DeleteApiKey deletes the specified API key.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/api-keys/-id-
func (s *ApiKeysService) DeleteApiKey(ctx context.Context, apiKeyID string) (*http.Response, error) {
	u := fmt.Sprintf(apikeysSinglePath, apiKeyID)
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

// ModifyApiKey updates an API key.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/api-keys/-id-
func (s *ApiKeysService) ModifyApiKey(ctx context.Context, apiKeyID string, request *ModifyApiKeyRequest) (*ApiKey, *http.Response, error) {
	u := fmt.Sprintf(apikeysSinglePath, apiKeyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	apiKey := new(ApiKey)
	resp, err := s.Client.Do(req, apiKey)
	if err != nil {
		return nil, resp, err
	}

	return apiKey, resp, nil
}

// Ensure ApiKey implementation satisfies the expected interfaces.
var (
	// ApiKey implements SubjectInterface
	_ service.SubjectInterface = &ApiKey{}
)

func (u *ApiKey) GetSubjectType() subjecttype.SubjectType { return subjecttype.ApiKey }
