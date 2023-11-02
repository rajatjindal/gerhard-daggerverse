package main

import (
	"context"
	"fmt"
)

const (
	// https://hub.docker.com/r/flyio/flyctl/tags
	version = "0.1.115"
)

type Flyio struct{}

// Example usage: "dagger call deploy --dir . --token "$YOUR_FLYIO_AUTH_TOKEN"
func (m *Flyio) Deploy(ctx context.Context, dir *Directory, token *Secret) (string, error) {
	return dag.Container().
		From(fmt.Sprintf("flyio/flyctl:v%s", version)).
		WithMountedFile("fly.toml", dir.File("fly.toml")).
		WithSecretVariable("FLY_API_TOKEN", token).
		WithExec([]string{"deploy"}).
		Stdout(ctx)
}
