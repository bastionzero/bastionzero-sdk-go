package sessionrecordings

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections/connectionstate"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	sessionRecordingsBasePath   = "api/v2/session-recordings"
	sessionRecordingsSinglePath = sessionRecordingsBasePath + "/%s"
)

// SessionRecordingsService handles communication with the session recording
// endpoints of the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Session-Recordings
type SessionRecordingsService client.Service

// SessionRecording describes a session recording of a specific connection
type SessionRecording struct {
	ConnectionID    string                          `json:"connectionId"`
	TimeCreated     types.Timestamp                 `json:"timeCreated"`
	ConnectionState connectionstate.ConnectionState `json:"connectionState"`
	TargetID        string                          `json:"targetId"`
	TargetType      targettype.TargetType           `json:"targetType"`
	TargetName      string                          `json:"targetName"`
	TargetUser      string                          `json:"targetUser"`
	InputRecorded   bool                            `json:"inputRecorded"`
	SubjectID       string                          `json:"subjectId"`
	Size            int                             `json:"size"`
}

// GetSessionRecordingFile fetches the session recording file for the specified connection ID.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/session-recordings/-connectionId-
func (s *SessionRecordingsService) GetSessionRecordingFile(ctx context.Context, connectionID string) (string, *http.Response, error) {
	u := fmt.Sprintf(sessionRecordingsSinglePath, connectionID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", nil, err
	}

	var file bytes.Buffer
	resp, err := s.Client.Do(req, &file)
	if err != nil {
		return "", resp, err
	}

	return file.String(), resp, nil
}

// DeleteSessionRecordingFile deletes the session recording file for the specified connection ID.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/session-recordings/-connectionId-
func (s *SessionRecordingsService) DeleteSessionRecordingFile(ctx context.Context, connectionID string) (*http.Response, error) {
	u := fmt.Sprintf(sessionRecordingsSinglePath, connectionID)
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

// ListSessionRecordings lists all session recordings.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/session-recordings
func (s *SessionRecordingsService) ListSessionRecordings(ctx context.Context) ([]SessionRecording, *http.Response, error) {
	u := sessionRecordingsBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	sessionRecordingList := new([]SessionRecording)
	resp, err := s.Client.Do(req, sessionRecordingList)
	if err != nil {
		return nil, resp, err
	}

	return *sessionRecordingList, resp, nil
}
