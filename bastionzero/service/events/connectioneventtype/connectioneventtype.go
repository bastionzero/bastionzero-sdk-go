package connectioneventtype

//go:generate go run github.com/lindell/string-enumer -t ConnectionEventType -o ./generated.go .

// ConnectionEventType represents the type of connection event
type ConnectionEventType string

const (
	// Created denotes that a connection was created
	Created ConnectionEventType = "Created"
	// Closed denotes that a connection was closed
	Closed ConnectionEventType = "Closed"
	// ClientConnect denotes that the client/daemon has connected
	ClientConnect ConnectionEventType = "ClientConnect"
	// ClientDisconnect denotes that the client/daemon has disconnected
	ClientDisconnect ConnectionEventType = "ClientDisconnect"
	// ShellConnect denotes that the shell interface has connected
	ShellConnect ConnectionEventType = "ShellConnect"
	// ShellDisconnect denotes that the shell interface has been disconnected
	ShellDisconnect ConnectionEventType = "ShellDisconnect"
)
