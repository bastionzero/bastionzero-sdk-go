package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	usersBasePath   = "api/v2/users"
	usersSinglePath = usersBasePath + "/%s"
)

// UsersService handles communication with the users endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Users
type UsersService client.Service

// User is a BastionZero user belonging to an organization
type User struct {
	ID             string           `json:"id"`
	OrganizationID string           `json:"organizationId"`
	FullName       string           `json:"fullName"`
	Email          string           `json:"email"`
	IsAdmin        bool             `json:"isAdmin"`
	TimeCreated    types.Timestamp  `json:"timeCreated"`
	LastLogin      *types.Timestamp `json:"lastLogin,omitempty"`
}

// ListUsers lists all users in your organization.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/users
func (s *UsersService) ListUsers(ctx context.Context) ([]User, *http.Response, error) {
	u := usersBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	userList := new([]User)
	resp, err := s.Client.Do(req, userList)
	if err != nil {
		return nil, resp, err
	}

	return *userList, resp, nil
}

// GetUser fetches the specified user by ID or email address.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/users/-id-, https://cloud.bastionzero.com/api/#get-/api/v2/users/-email-
func (s *UsersService) GetUser(ctx context.Context, userIDOrEmail string) (*User, *http.Response, error) {
	u := fmt.Sprintf(usersSinglePath, userIDOrEmail)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.Client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}
