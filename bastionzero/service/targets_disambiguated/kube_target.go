package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
)

type KubeTarget struct {
	Target

	AllowedClusterUsers  []string                     `json:"allowedClusterUsers"`
	AllowedClusterGroups []string                     `json:"allowedClusterGroups"`
	ValidClusterUsers    []string                     `json:"validClusterUsers"`
	Connections          []connections.KubeConnection `json:"connections"`
}
