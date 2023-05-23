package oktapublickeys

import (
	"context"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	oktaPublicKeysBasePath = "api/v2/okta-public-keys"
)

// OktaPublicKeysService handles communication with the okta public keys
// endpoints of the BastionZero API.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#tag--Okta-Public-Keys
type OktaPublicKeysService client.Service

// OktaPublicKey represents a public key that can be used to verify a JWT client
// assertion signed by BastionZero.
type OktaPublicKey struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	N   string `json:"n"`
}

// ListOktaPublicKeysResponse is the response returned when querying for Okta
// public keys
type ListOktaPublicKeysResponse struct {
	Keys []OktaPublicKey `json:"keys"`
}

// ListOktaPublicKeys lists Okta public keys.  The list of the returned public
// keys can be used to verify that the JWT client assertion was signed by
// BastionZero. The JWT assertion is used by BastionZero in order to acquire
// access tokens on behalf of a client. These access tokens are then used by
// BastionZero to fetch Groups.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/okta-public-keys
func (s *OktaPublicKeysService) ListOktaPublicKeys(ctx context.Context) (*ListOktaPublicKeysResponse, *http.Response, error) {
	u := oktaPublicKeysBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	listOktaPublicKeysResponse := new(ListOktaPublicKeysResponse)
	resp, err := s.Client.Do(req, listOktaPublicKeysResponse)
	if err != nil {
		return nil, resp, err
	}

	return listOktaPublicKeysResponse, resp, nil
}
