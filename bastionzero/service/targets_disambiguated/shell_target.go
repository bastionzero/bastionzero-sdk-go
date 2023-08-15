package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

type ShellTarget struct {
	Target

	DynamicAccess      bool                          `json:"dynamicAccess"`
	AllowedTargetUsers []policies.TargetUser         `json:"allowedTargetUsers"`
	Connections        []connections.ShellConnection `json:"connections"`
}
