package autodiscoveryscripts

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/autodiscoveryscripts/targetnameoption"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

// BzeroBashAutodiscoveryOptions specifies the required query-string parameters
// in order to get a Linux bash autodiscovery script
type BzeroBashAutodiscoveryOptions struct {
	// TargetNameOption specifies the target name schema option. Required.
	TargetNameOption targetnameoption.TargetNameOption `url:"targetNameOption"`

	// EnvironmentID is the unique ID for the environment the target should
	// associate with. Required.
	EnvironmentID string `url:"environmentId"`
}

// WindowsPowershellAutodiscoveryOptions specifies the required query-string parameters
// in order to get a Windows Powershell autodiscovery script
type WindowsPowershellAutodiscoveryOptions struct {
	// TargetNameOption specifies the target name schema option. Required.
	TargetNameOption targetnameoption.TargetNameOption `url:"targetNameOption"`

	// EnvironmentID is the unique ID for the environment the target should
	// associate with. Required.
	EnvironmentID string `url:"environmentId"`
}

type BzeroBashAutodiscoveryScript struct {
	Script string `json:"autodiscoveryScript"`
}

type WindowsPowershellAutodiscoveryScript struct {
	Script string `json:"autodiscoveryScript"`
}

// GetBzeroBashAutodiscoveryScript gets a bash script that can be used to
// install the latest production BastionZero agent (bzero) on your Linux targets.
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

// GetBzeroPowershellAutodiscoveryScript gets a powershell script that can be used to
// install the latest production BastionZero agent (bzero) on your Windows targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/autodiscovery-scripts/windows/powershell
func (s *AutodiscoveryScriptsService) GetBzeroPowershellAutodiscoveryScript(ctx context.Context, opts *WindowsPowershellAutodiscoveryOptions) (*WindowsPowershellAutodiscoveryScript, *http.Response, error) {
	u := autodiscoveryScriptsBasePath + "/windows/powershell"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	script := new(WindowsPowershellAutodiscoveryScript)
	resp, err := s.Client.Do(req, script)
	if err != nil {
		return nil, resp, err
	}

	return script, resp, nil
}

// BzeroAnsibleAutodiscoveryOptions specifies the required query-string
// parameters in order to get a Bzero Ansible autodiscovery playbook
type BzeroAnsibleAutodiscoveryOptions struct {
	// EnvironmentID is the unique ID for the environment the target should
	// associate with. Required.
	EnvironmentID string `url:"environmentId"`
}

type BzeroAnsibleAutodiscoveryPlaybook struct {
	Playbook string `json:"autodiscoveryScript"`
}

// GetBzeroAnsibleAutodiscoveryPlaybook gets an Ansible playbook that can be
// used to install the latest production BastionZero agent (bzero) on your
// targets.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/autodiscovery-scripts/bzero/ansible
func (s *AutodiscoveryScriptsService) GetBzeroAnsibleAutodiscoveryPlaybook(ctx context.Context, opts *BzeroAnsibleAutodiscoveryOptions) (*BzeroAnsibleAutodiscoveryPlaybook, *http.Response, error) {
	u := autodiscoveryScriptsBasePath + "/bzero/ansible"
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	playbook := new(BzeroAnsibleAutodiscoveryPlaybook)
	resp, err := s.Client.Do(req, playbook)
	if err != nil {
		return nil, resp, err
	}

	return playbook, resp, nil
}
