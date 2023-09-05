package targets_disambiguated

import (
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/connections"
	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies"
)

type DatabaseTarget struct {
	Target

	ProxyAgentId       string                     `json:"proxyAgentId"`
	ProxyAgentName     string                     `json:"proxyAgentName"`
	RemoteHost         string                     `json:"remoteHost"`
	RemotePort         Port                       `json:"remotePort"`
	LocalHost          string                     `json:"localHost"`
	LocalPort          *Port                      `json:"localPort"`
	SplitCert          bool                       `json:"splitCert"`
	DatabaseType       string                     `json:"databaseType"`
	AllowedTargetUsers []policies.TargetUser      `json:"allowedTargetUsers"`
	Connections        []connections.DbConnection `json:"connections"`
}
