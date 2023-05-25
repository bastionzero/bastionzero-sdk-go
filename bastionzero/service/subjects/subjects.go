package subjects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/subjects/roletype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/subjecttype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	subjectsBasePath   = "api/v2/subjects"
	subjectsSinglePath = subjectsBasePath + "/%s"
)

// SubjectsService handles communication with the subjects endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Subjects
type SubjectsService client.Service

// ModifySubjectRoleRequest is used to modify a subject's role
type ModifySubjectRoleRequest struct {
	Role roletype.RoleType `json:"role"`
}

// Subject is the BastionZero entity that can interact with the BastionZero API
type Subject struct {
	ID             string                  `json:"id"`
	OrganizationID string                  `json:"organizationId"`
	Email          string                  `json:"email"`
	IsAdmin        bool                    `json:"isAdmin"`
	LastLogin      *types.Timestamp        `json:"lastLogin"`
	TimeCreated    types.Timestamp         `json:"timeCreated"`
	Type           subjecttype.SubjectType `json:"type"`
}

// Me fetches your subject information (current subject).
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/subjects/me
func (s *SubjectsService) Me(ctx context.Context) (*Subject, *http.Response, error) {
	u := subjectsBasePath + "/me"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	subject := new(Subject)
	resp, err := s.Client.Do(req, subject)
	if err != nil {
		return nil, resp, err
	}

	return subject, resp, nil
}

// GetSubject fetches the specified subject by ID or email address.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/subjects/-id-, https://cloud.bastionzero.com/api/#get-/api/v2/subjects/-email-
func (s *SubjectsService) GetSubject(ctx context.Context, subjectIDOrEmail string) (*Subject, *http.Response, error) {
	u := fmt.Sprintf(subjectsSinglePath, subjectIDOrEmail)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	subject := new(Subject)
	resp, err := s.Client.Do(req, subject)
	if err != nil {
		return nil, resp, err
	}

	return subject, resp, nil
}

// ModifySubjectRole updates the specified subject's role.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/subjects/-id-
func (s *SubjectsService) ModifySubjectRole(ctx context.Context, id string, request *ModifySubjectRoleRequest) (*http.Response, error) {
	u := fmt.Sprintf(subjectsSinglePath, id)
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

// ListSubjects lists all subjects.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/subjects
func (s *SubjectsService) ListSubjects(ctx context.Context) ([]Subject, *http.Response, error) {
	u := subjectsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	subjectList := new([]Subject)
	resp, err := s.Client.Do(req, subjectList)
	if err != nil {
		return nil, resp, err
	}

	return *subjectList, resp, nil
}

// CloseSubjectConnections closes all connections of the specified subject.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/subjects/-id-/close-connections
func (s *SubjectsService) CloseSubjectConnections(ctx context.Context, id string) (*http.Response, error) {
	u := fmt.Sprintf(subjectsSinglePath+"/close-connections", id)
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
