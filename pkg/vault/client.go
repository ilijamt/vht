package vault

import (
	"github.com/hashicorp/vault/api"
	"github.com/ilijamt/vht/pkg/vault/config"
	"github.com/pkg/errors"
	"os"
	"time"
)

func Client() (client *api.Client, err error) {
	cfg := api.DefaultConfig()
	cfg.HttpClient.Timeout = 10 * time.Second

	if err := cfg.ReadEnvironment(); err != nil {
		return nil, errors.Wrap(err, "failed to read environment")
	}

	// Build the client
	client, err = api.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	// Turn off retries on the CLI
	if os.Getenv(api.EnvVaultMaxRetries) == "" {
		client.SetMaxRetries(0)
	}

	// Get the token if it came in from the environment
	token := client.Token()

	// If we don't have a token, check the token helper
	if token == "" {
		helper, err := config.DefaultTokenHelper()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get token helper")
		}
		token, err = helper.Get()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get token from token helper")
		}
	}

	// Set the token
	if token != "" {
		client.SetToken(token)
	}

	return client, err
}
