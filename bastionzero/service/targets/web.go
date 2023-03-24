package targets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/types/targettype"
)

const (
	webBasePath   = targetsBasePath + "/web"
	webSinglePath = webBasePath + "/%s"
)

// ModifyWebTargetRequest is used to modify a Web target
type ModifyWebTargetRequest struct {
	TargetName    *string `json:"targetName,omitempty"`
	ProxyTargetID *string `json:"proxyTargetId,omitempty"`
	RemoteHost    *string `json:"remoteHost,omitempty"`
	RemotePort    *Port   `json:"remotePort,omitempty"`
	LocalPort     *Port   `json:"localPort,omitempty"`
	LocalHost     *string `json:"localHost,omitempty"`
	EnvironmentID *string `json:"environmentId,omitempty"`
}

// WebTarget is a virtual target that provides HTTP(S) access to a remote web
// server. The connection is proxied by a Bzero or Cluster target.
type WebTarget struct {
	*VirtualTarget
}

// ListWebTargets lists all Web targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/web
func (s *TargetsService) ListWebTargets(ctx context.Context) ([]WebTarget, *http.Response, error) {
	u := webBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetList := new([]WebTarget)
	resp, err := s.Client.Do(req, targetList)
	if err != nil {
		return nil, resp, err
	}

	return *targetList, resp, nil
}

// GetWebTarget fetches the specified Web target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/targets/web/-id-
func (s *TargetsService) GetWebTarget(ctx context.Context, targetID string) (*WebTarget, *http.Response, error) {
	u := fmt.Sprintf(webSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	target := new(WebTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}

// ModifyWebTarget updates a Web target.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/targets/web/-id-
func (s *TargetsService) ModifyWebTarget(ctx context.Context, targetID string, request *ModifyWebTargetRequest) (*WebTarget, *http.Response, error) {
	u := fmt.Sprintf(webSinglePath, targetID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, request)
	if err != nil {
		return nil, nil, err
	}

	target := new(WebTarget)
	resp, err := s.Client.Do(req, target)
	if err != nil {
		return nil, resp, err
	}

	return target, resp, nil
}

// Ensure WebTarget implementation satisfies the expected interfaces.
var (
	// WebTarget implements VirtualTargetInterface
	_ VirtualTargetInterface = &WebTarget{}
)

func (t *WebTarget) GetTargetType() targettype.TargetType { return targettype.Web }
