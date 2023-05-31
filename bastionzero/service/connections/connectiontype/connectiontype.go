package connectiontype

//go:generate go run github.com/lindell/string-enumer -t ConnectionType -o ./generated.go .

// ConnectionType represents the type of a connection
type ConnectionType string

const (
	// Shell represents a shell connection
	Shell ConnectionType = "Shell"
	// Dynamic represents a connection to a dynamic access target (DAT)
	Dynamic ConnectionType = "Dynamic"
	// Kube represents a Kube connection
	Kube ConnectionType = "Kube"
	// Web represents a Web connection
	Web ConnectionType = "Web"
	// Db represents a Db connection
	Db ConnectionType = "Db"
	// Ssh represents a ssh connection
	Ssh ConnectionType = "Ssh"
)
