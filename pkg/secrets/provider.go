package secrets

import "context"

// Provider secrets provider interface
type Provider interface {
	ResolveSecrets(ctx context.Context, rawValues map[string]string, envs []string) ([]string, error)
}
