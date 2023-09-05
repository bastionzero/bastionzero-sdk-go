package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
)

type RDPTarget struct {
	Target

	RemoteHost  string                      `json:"remoteHost"`
	RemotePort  Port                        `json:"remotePort"`
	Connections []connections.RDPConnection `json:"connections"`
}
