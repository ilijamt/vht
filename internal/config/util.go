package config

import "github.com/ilijamt/vht/internal/token"

// DefaultTokenHelper returns the token helper that is configured for Vault.
func DefaultTokenHelper() (token.TokenHelper, error) {
	config, err := LoadConfig("")
	if err != nil {
		return nil, err
	}

	path := config.TokenHelper
	if path == "" {
		return token.NewInternalTokenHelper()
	}

	path, err = token.ExternalTokenHelperPath(path)
	if err != nil {
		return nil, err
	}
	return &token.ExternalTokenHelper{BinaryPath: path}, nil
}
