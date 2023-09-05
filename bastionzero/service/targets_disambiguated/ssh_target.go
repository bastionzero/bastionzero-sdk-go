package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

type SSHTarget struct {
	Target

	AllowedTargetUsers []policies.TargetUser       `json:"allowedTargetUsers"`
	Connections        []connections.SSHConnection `json:"connections"`
}
