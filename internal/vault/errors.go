package vault

const (
	// ErrMissingVaultClient means that we have not provided a valid vault client
	ErrMissingVaultClient            = "missing vault client"
	ErrMissingVaultAddrOrCredentials = "missing vault address or token"
)
