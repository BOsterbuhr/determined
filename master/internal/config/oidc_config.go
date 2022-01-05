package config

import (
	"net/url"
)

// OIDCConfig holds the parameters for the OIDC provider.
type OIDCConfig struct {
	Enabled         bool   `json:"enabled"`
	Provider        string `json:"provider"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	IDPSSOURL       string `json:"idp_sso_url"`
	IDPRecipientURL string `json:"idp_recipient_url"`
}

// Validate implements the check.Validatable interface.
func (c OIDCConfig) Validate() []error {
	if !c.Enabled {
		return nil
	}

	_, err := url.Parse(c.IDPRecipientURL)
	return []error{err}
}
