package targetnameoption

//go:generate go run github.com/lindell/string-enumer -t TargetNameOption -o ./generated.go .

// TargetNameOption represents the target name schema option to use during
// autodiscovery
type TargetNameOption string

const (
	// Timestamp indicates that the current timestamp should be used as the
	// target's name
	Timestamp TargetNameOption = "Timestamp"
	// DigitalOceanMetadata indicates that the DigitalOcean droplet's hostname
	// should be used as the target's name
	DigitalOceanMetadata TargetNameOption = "DigitalOceanMetadata"
	// AwsEc2Metadata indicates that the AWS instance's ID should be used as the
	// target's name
	AwsEc2Metadata TargetNameOption = "AwsEc2Metadata"
	// BashHostName indicates that the machine's hostname should be used as the
	// target's name
	BashHostName TargetNameOption = "BashHostName"
)
