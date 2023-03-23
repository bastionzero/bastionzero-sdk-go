package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	sessionRecordingBasePath   = policiesBasePath + "/session-recording"
	sessionRecordingSinglePath = sessionRecordingBasePath + "/%s"
)

// SessionRecordingPolicy represents a session recording policy. Session
// recording policies govern whether users' I/O during shell connections are
// recorded.
type SessionRecordingPolicy struct {
	*Policy
	RecordInput *bool `json:"recordInput,omitempty"`
}

// ListSessionRecordingPolicies lists all session recording policies.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/session-recording
func (s *PoliciesService) ListSessionRecordingPolicies(ctx context.Context, opts *ListPolicyOptions) ([]SessionRecordingPolicy, *http.Response, error) {
	u := sessionRecordingBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]SessionRecordingPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateSessionRecordingPolicy creates a new session recording policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/session-recording
func (s *PoliciesService) CreateSessionRecordingPolicy(ctx context.Context, policy *SessionRecordingPolicy) (*SessionRecordingPolicy, *http.Response, error) {
	u := sessionRecordingBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(SessionRecordingPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetSessionRecordingPolicy fetches the specified session recording policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/session-recording/-id-
func (s *PoliciesService) GetSessionRecordingPolicy(ctx context.Context, policyID string) (*SessionRecordingPolicy, *http.Response, error) {
	u := fmt.Sprintf(sessionRecordingSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(SessionRecordingPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteSessionRecordingPolicy deletes the specified session recording policy.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/session-recording/-id-
func (s *PoliciesService) DeleteSessionRecordingPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(sessionRecordingSinglePath, policyID)
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

// ModifySessionRecordingPolicy updates a session recording policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/session-recording/-id-
func (s *PoliciesService) ModifySessionRecordingPolicy(ctx context.Context, policyID string, policy *SessionRecordingPolicy) (*SessionRecordingPolicy, *http.Response, error) {
	u := fmt.Sprintf(sessionRecordingSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(SessionRecordingPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure SessionRecordingPolicy implementation satisfies the expected
// interfaces.
var (
	// SessionRecordingPolicy implements PolicyInterface
	_ PolicyInterface = &SessionRecordingPolicy{}
)

func (p *SessionRecordingPolicy) GetID() string                    { return p.ID }
func (p *SessionRecordingPolicy) GetTimeExpires() *types.Timestamp { return p.TimeExpires }
func (p *SessionRecordingPolicy) GetName() string                  { return p.Name }
func (p *SessionRecordingPolicy) GetDescription() string {
	if p.Description == nil {
		return ""
	}
	return *p.Description
}
func (p *SessionRecordingPolicy) GetSubjects() []Subject {
	if p.Subjects == nil {
		return []Subject{}
	}
	return *p.Subjects
}
func (p *SessionRecordingPolicy) GetGroups() []Group {
	if p.Groups == nil {
		return []Group{}
	}
	return *p.Groups
}
func (p *SessionRecordingPolicy) GetPolicyType() policytype.PolicyType {
	return policytype.SessionRecording
}

func (p *SessionRecordingPolicy) GetRecordInput() bool {
	if p.RecordInput == nil {
		return false
	}
	return *p.RecordInput
}
