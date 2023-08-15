package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

type FileTransferTarget struct {
	Target

	AllowedTargetUsers []policies.TargetUser `json:"allowedTargetUsers"`
}
