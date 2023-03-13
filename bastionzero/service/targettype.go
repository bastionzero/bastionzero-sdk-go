package service

//go:generate go run github.com/dmarkham/enumer -type=TargetType -json

// TargetType represents the type of target
type TargetType int

const (
	// Bzero represents a Bzero target
	Bzero TargetType = iota
	// Cluster represents a Kubernetes target
	Cluster

	// DynamicAccessConfig represents a DAC target
	DynamicAccessConfig
	// Web represents a Web target
	Web
	// Db represents a Db target
	Db
)
