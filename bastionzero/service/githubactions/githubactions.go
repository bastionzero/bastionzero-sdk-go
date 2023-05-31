package githubactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	githubactionsBasePath   = "api/v2/github-actions"
	githubactionsSinglePath = githubactionsBasePath + "/%s"
)

// GitHubActionsService handles communication with the github actions endpoints
// of the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Authorized-Github-Actions
type GitHubActionsService client.Service

// CreateAuthorizedGitHubActionRequest is used to create a new authorized GitHub action
type CreateAuthorizedGitHubActionRequest struct {
	GitHubActionId string `json:"githubActionId"`
}

// AuthorizedGitHubAction is a GitHub action that can approve Just in Time (JIT)
// access for users and service accounts
type AuthorizedGitHubAction struct {
	ID             string          `json:"id"`
	OrganizationId string          `json:"organizationId"`
	TimeCreated    types.Timestamp `json:"timeCreated"`
	CreatedBy      string          `json:"createdBy"`
	GitHubActionId string          `json:"githubActionId"`
}

// ListAuthorizedGitHubActions lists all authorized GitHub actions.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/github-actions
func (s *GitHubActionsService) ListAuthorizedGitHubActions(ctx context.Context) ([]AuthorizedGitHubAction, *http.Response, error) {
	u := githubactionsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	authorizedGitHubActionList := new([]AuthorizedGitHubAction)
	resp, err := s.Client.Do(req, authorizedGitHubActionList)
	if err != nil {
		return nil, resp, err
	}

	return *authorizedGitHubActionList, resp, nil
}

// CreateAuthorizedGitHubAction creates a new authorized GitHub action.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/github-actions
func (s *GitHubActionsService) CreateAuthorizedGitHubAction(ctx context.Context, request *CreateAuthorizedGitHubActionRequest) (*AuthorizedGitHubAction, *http.Response, error) {
	u := githubactionsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	authGitHubAction := new(AuthorizedGitHubAction)
	resp, err := s.Client.Do(req, authGitHubAction)
	if err != nil {
		return nil, resp, err
	}

	return authGitHubAction, resp, nil
}

// GetAuthorizedGitHubAction fetches the specified authorized GitHub action.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/github-actions/-id-
func (s *GitHubActionsService) GetAuthorizedGitHubAction(ctx context.Context, authGitHubActionId string) (*AuthorizedGitHubAction, *http.Response, error) {
	u := fmt.Sprintf(githubactionsSinglePath, authGitHubActionId)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	authGitHubAction := new(AuthorizedGitHubAction)
	resp, err := s.Client.Do(req, authGitHubAction)
	if err != nil {
		return nil, resp, err
	}

	return authGitHubAction, resp, nil
}

// DeleteAuthorizedGitHubAction deletes the specified authorized GitHub action.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/github-actions/-id-
func (s *GitHubActionsService) DeleteAuthorizedGitHubAction(ctx context.Context, authGitHubActionId string) (*http.Response, error) {
	u := fmt.Sprintf(githubactionsSinglePath, authGitHubActionId)
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
