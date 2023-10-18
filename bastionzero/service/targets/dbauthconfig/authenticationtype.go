package dbauthconfig

// These constants represent the supported values for the AuthenticationType field in DatabaseAuthenticationConfig.
const (
	Default                 string = "Default"
	SplitCert               string = "SplitCert"
	ServiceAccountInjection string = "ServiceAccountInjection"
)
