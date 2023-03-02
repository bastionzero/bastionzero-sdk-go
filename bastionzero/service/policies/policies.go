package policies

import (
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	policiesBasePath = "api/v2/policies"
)

// PoliciesService handles communication with the policies endpoints of the
// BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Policies
type PoliciesService client.Service
