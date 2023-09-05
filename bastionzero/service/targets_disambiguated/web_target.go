package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
)

type WebTarget struct {
	Target

	ProxyAgentId   string                      `json:"proxyAgentId"`
	ProxyAgentName string                      `json:"proxyAgentName"`
	RemoteHost     string                      `json:"remoteHost"`
	RemotePort     Port                        `json:"remotePort"`
	LocalHost      string                      `json:"localHost"`
	LocalPort      *Port                       `json:"localPort"`
	Connections    []connections.WebConnection `json:"connections"`
}
