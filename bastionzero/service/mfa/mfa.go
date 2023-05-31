package mfa

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	mfaBasePath = "api/v2/mfa"
)

// MFAService handles communication with the mfa endpoints of the BastionZero
// API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--MFA
type MFAService client.Service

// ResetMFASecretRequest is used to reset your MFA secret
type ResetMFASecretRequest struct {
	ForceSetup bool `json:"forceSetup"`
}

// ResetMFASecretResponse is the response returned if one's MFA secret has been reset
type ResetMFASecretResponse struct {
	MFASecretUrl string `json:"mfaSecretUrl"`
}

// ClearMFASecretRequest is used to clear the MFA secret of a specific user
type ClearMFASecretRequest struct {
	UserID string `json:"userId"`
}

// EnableMFARequest is used to enable MFA for a specific user
type EnableMFARequest struct {
	UserID string `json:"userId"`
}

// DisableMFARequest is used to disable MFA for a specific user
type DisableMFARequest struct {
	UserID string `json:"userId"`
}

// MFAStatus describes a user's MFA status
type MFAStatus struct {
	Enabled            bool             `json:"enabled"`
	Verified           bool             `json:"verified"`
	SessionVerified    *bool            `json:"sessionVerified"`
	GracePeriodEndTime *types.Timestamp `json:"gracePeriodEndTime"`
}

// ResetMFASecret resets your MFA secret. After this operation is completed, you
// will have to scan or manually enter the new secret into your authenticator
// app the next time you log in.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/mfa/reset
func (s *MFAService) ResetMFASecret(ctx context.Context, request *ResetMFASecretRequest) (*ResetMFASecretResponse, *http.Response, error) {
	u := mfaBasePath + "/reset"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	resetMFAResponse := new(ResetMFASecretResponse)
	resp, err := s.Client.Do(req, resetMFAResponse)
	if err != nil {
		return nil, resp, err
	}

	return resetMFAResponse, resp, nil
}

// RotateServiceAccountMFASecret rotates the MFA secret of a service account.
// After this operation is completed, all existing sessions of the specified
// service account will become unauthorized. Returns the new MFA shared secret
// of the specified service account in Base32 encoding.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/mfa/rotate/-id-
func (s *MFAService) RotateServiceAccountMFASecret(ctx context.Context, serviceAccountID string) (string, *http.Response, error) {
	u := mfaBasePath + fmt.Sprintf("/rotate/%s", serviceAccountID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, nil)
	if err != nil {
		return "", nil, err
	}

	var mfaSecret string
	resp, err := s.Client.Do(req, &mfaSecret)
	if err != nil {
		return "", resp, err
	}

	return mfaSecret, resp, nil
}

// ClearMFASecret clears the MFA secret of a specific user. After this operation is
// completed, the user's MFA secret will be reset on the next attempt to log in.
// The user will have to scan or manually enter the new secret into an
// authenticator app.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/mfa/clear
func (s *MFAService) ClearMFASecret(ctx context.Context, request *ClearMFASecretRequest) (*http.Response, error) {
	u := mfaBasePath + "/clear"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// EnableMFA turns on MFA for a specific user. After MFA is enabled, on their
// next attempt to log in, the user will be prompted to scan or manually enter
// their MFA secret into an authenticator app.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/mfa/setup
func (s *MFAService) EnableMFA(ctx context.Context, request *EnableMFARequest) (*http.Response, error) {
	u := mfaBasePath + "/setup"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DisableMFA turns off MFA for a specific user.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/mfa/disable
func (s *MFAService) DisableMFA(ctx context.Context, request *DisableMFARequest) (*http.Response, error) {
	u := mfaBasePath + "/disable"
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, request)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetUserMFAStatus fetches the MFA status for a specific user.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/mfa/-userId-
func (s *MFAService) GetUserMFAStatus(ctx context.Context, userID string) (*MFAStatus, *http.Response, error) {
	u := mfaBasePath + fmt.Sprintf("/%s", userID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	mfaStatus := new(MFAStatus)
	resp, err := s.Client.Do(req, mfaStatus)
	if err != nil {
		return nil, resp, err
	}

	return mfaStatus, resp, nil
}

// GetMFAStatus fetches your MFA status (current subject).
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/mfa/me
func (s *MFAService) GetMFAStatus(ctx context.Context) (*MFAStatus, *http.Response, error) {
	u := mfaBasePath + "/me"
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	mfaStatus := new(MFAStatus)
	resp, err := s.Client.Do(req, mfaStatus)
	if err != nil {
		return nil, resp, err
	}

	return mfaStatus, resp, nil
}
