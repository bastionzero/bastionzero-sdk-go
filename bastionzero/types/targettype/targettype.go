package targettype

//go:generate go run github.com/lindell/string-enumer -t TargetType -o ./generated.go .

// TODO-Yuval: Move this entire package/folder to service/targets batched in a
// breaking change release

// TargetType represents the type of target
type TargetType string

const (
	// Bzero represents a Bzero target
	Bzero TargetType = "Bzero"
	// Cluster represents a Kubernetes target
	Cluster TargetType = "Cluster"

	// DynamicAccessConfig represents a DAC target
	DynamicAccessConfig TargetType = "DynamicAccessConfig"
	// Web represents a Web target
	Web TargetType = "Web"
	// Db represents a Db target
	Db TargetType = "Db"
)
