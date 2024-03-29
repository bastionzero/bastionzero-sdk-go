package verbtype

//go:generate go run github.com/lindell/string-enumer -t VerbType -o ./generated.go .

// VerbType represents the type of target connect verb
type VerbType string

const (
	// Shell represents the ability to make a Shell connection
	Shell VerbType = "Shell"
	// FileTransfer represents the ability to upload/download files
	FileTransfer VerbType = "FileTransfer"
	// Tunnel represents the ability to make an SSH tunnel
	Tunnel VerbType = "Tunnel"
	// RDP represents the ability to connect using the Remote Desktop Protocol
	RDP VerbType = "RDP"
	// SQLServer represents the ability to connect to a Windows SQL Server
	SQLServer VerbType = "SQLServer"
)
