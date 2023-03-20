package dacstatus

//go:generate go run github.com/lindell/string-enumer -t DynamicAccessConfigurationStatus -o ./generated.go .

// DynamicAccessConfigurationStatus represents the status of a dynamic access
// configuration
type DynamicAccessConfigurationStatus string

const (
	// DACOffline indicates the health webhook of the DAC is responding
	// unhealthy
	DACOffline DynamicAccessConfigurationStatus = "Offline"
	// DACOnline indicates the health webhook of the DAC is responding healthy
	DACOnline DynamicAccessConfigurationStatus = "Online"
)
