package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/subjects/roletype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
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

// ModifyUserRoleRequest is used to modify a user's role
type ModifyUserRoleRequest struct {
	Role roletype.RoleType `json:"role"`
}

// User is a BastionZero user belonging to an organization
type User struct {
	service.Subject

	OrganizationID string           `json:"organizationId"`
	FullName       string           `json:"fullName"`
	Email          string           `json:"email"`
	IsAdmin        bool             `json:"isAdmin"`
	TimeCreated    types.Timestamp  `json:"timeCreated"`
	LastLogin      *types.Timestamp `json:"lastLogin"`
}

// Me fetches your user information (current subject).
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/users/me
func (s *UsersService) Me(ctx context.Context) (*User, *http.Response, error) {
	u := usersBasePath + "/me"
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

// DeleteUser deletes the specified user.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/users/-id-
func (s *UsersService) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	u := fmt.Sprintf(usersSinglePath, id)
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

// ModifyUserRole updates the specified user's role.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/users/-id-
func (s *UsersService) ModifyUserRole(ctx context.Context, id string, request *ModifyUserRoleRequest) (*http.Response, error) {
	u := fmt.Sprintf(usersSinglePath, id)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
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

// CloseUserConnections closes all connections of the specified user.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/users/-id-/close-connections
func (s *UsersService) CloseUserConnections(ctx context.Context, id string) (*http.Response, error) {
	u := fmt.Sprintf(usersSinglePath+"/close-connections", id)
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

// Ensure User implementation satisfies the expected interfaces.
var (
	// User implements SubjectInterface
	_ service.SubjectInterface = &User{}
)

func (u *User) GetSubjectType() subjecttype.SubjectType { return subjecttype.User }
