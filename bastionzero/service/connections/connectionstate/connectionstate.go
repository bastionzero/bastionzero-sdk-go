package connectionstate

//go:generate go run github.com/lindell/string-enumer -t ConnectionState -o ./generated.go .

// ConnectionState represents the state of a connection
type ConnectionState string

const (
	// Open denotes that a connection is open
	Open ConnectionState = "Open"
	// Closed denotes that a connection is closed
	Closed ConnectionState = "Closed"
	// Error denotes that an error occurred when establishing the connection
	Error ConnectionState = "Error"
	// Pending denotes that a connection is still in the process of being opened
	Pending ConnectionState = "Pending"
)
