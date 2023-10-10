package dbauthconfig

// DatabaseAuthenticationConfig defines a database authentication configuration supported
// by BastionZero. When using a non-null DatabaseAuthenticationConfig in a request, it
// is recommended that the supported configurations are retrieved from a GET request to
// /api/v2/targets/database/supported-database-configs and then one of the returned
// configurations is used in any subsequent create or update request as needed.
type DatabaseAuthenticationConfig struct {
	AuthenticationType   *string `json:"authenticationType"`
	CloudServiceProvider *string `json:"cloudServiceProvider"`
	Database             *string `json:"database"`
	Label                *string `json:"label"`
}
