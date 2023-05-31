package autodiscoveryscripts

import (
	"bytes"
	"context"
	"net/http"
)

// GetContainerBashAutodiscoveryScript gets a bash script that can be used to
// install the latest production BastionZero agent (bzero) on containers.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/autodiscovery-scripts/container
func (s *AutodiscoveryScriptsService) GetContainerBashAutodiscoveryScript(ctx context.Context) (string, *http.Response, error) {
	u := autodiscoveryScriptsBasePath + "/container"

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", nil, err
	}

	var script bytes.Buffer
	resp, err := s.Client.Do(req, &script)
	if err != nil {
		return "", resp, err
	}

	return script.String(), resp, nil
}
