package datstate

//go:generate go run github.com/lindell/string-enumer -t DATState -o ./generated.go .

// DATState represents the state of a dynamic access target (DAT)
type DATState string

const (
	// Starting denotes that the provisioning server's start webhook is
	// currently being triggered
	Starting DATState = "Starting"
	// Started denotes that the start webhook was successfully triggered
	Started DATState = "Started"
	// StartError denotes that an error occurred when triggering the start
	// webhook
	StartError DATState = "StartError"
	// Stopping denotes that the provisioning server's stop webhook is currently
	// being triggered
	Stopping DATState = "Stopping"
	// Stopped denotes that the stop webhook was successfully triggered
	Stopped DATState = "Stopped"
	// StopError denotes that an error occurred when triggering the stop webhook
	StopError DATState = "StopError"
)
