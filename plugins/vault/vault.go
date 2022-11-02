package vault

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func Executable_vault() schema.Executable {
	return schema.Executable{
		Runs:      []string{"vault"},
		Name:      "HashiCorp Vault CLI",
		DocsURL:   sdk.URL("https://developer.hashicorp.com/vault/docs/commands"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Credentials: []schema.CredentialType{
			AuthToken(),
		},
	}
}
