package autodiscoveryscripts

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/autodiscoveryscripts/targetnameoption"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

// BzeroBashAutodiscoveryOptions specifies the required query-string parameters
// in order to get a Bzero bash autodiscovery script
type BzeroBashAutodiscoveryOptions struct {
	// TargetNameOption specifies the target name schema option. Required.
	TargetNameOption targetnameoption.TargetNameOption `url:"targetNameOption"`

	// EnvironmentID is the unique ID for the environment the target should
	// associate with. Required.
	EnvironmentID string `url:"environmentId"`
}

type BzeroBashAutodiscoveryScript struct {
	Script string `json:"autodiscoveryScript"`
}

// GetBzeroBashAutodiscoveryScript gets a bash script that can be used to
// install the latest production BastionZero agent (bzero) on your targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/autodiscovery-scripts/bzero/bash
func (s *AutodiscoveryScriptsService) GetBzeroBashAutodiscoveryScript(ctx context.Context, opts *BzeroBashAutodiscoveryOptions) (*BzeroBashAutodiscoveryScript, *http.Response, error) {
	u := autodiscoveryScriptsBasePath + "/bzero/bash"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	script := new(BzeroBashAutodiscoveryScript)
	resp, err := s.Client.Do(req, script)
	if err != nil {
		return nil, resp, err
	}

	return script, resp, nil
}
