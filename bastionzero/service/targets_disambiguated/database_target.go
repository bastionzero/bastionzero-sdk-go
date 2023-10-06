package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/targets"
)

type DatabaseTarget struct {
	Target

	ProxyAgentId                 string                               `json:"proxyAgentId"`
	ProxyAgentName               string                               `json:"proxyAgentName"`
	RemoteHost                   string                               `json:"remoteHost"`
	RemotePort                   Port                                 `json:"remotePort"`
	LocalHost                    string                               `json:"localHost"`
	LocalPort                    *Port                                `json:"localPort"`
	AllowedTargetUsers           []policies.TargetUser                `json:"allowedTargetUsers"`
	Connections                  []connections.DbConnection           `json:"connections"`
	DatabaseAuthenticationConfig targets.DatabaseAuthenticationConfig `json:"databaseAuthenticationConfig"`
}
