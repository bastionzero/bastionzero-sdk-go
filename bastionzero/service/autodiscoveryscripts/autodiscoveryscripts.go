package autodiscoveryscripts

import "github.com/bastionzero/bastionzero-sdk-go/internal/client"

const (
	autodiscoveryScriptsBasePath = "api/v2/autodiscovery-scripts"
)

// AutodiscoveryScriptsService handles communication with the
// autodiscovery-scripts endpoints of the BastionZero API.
type AutodiscoveryScriptsService client.Service
