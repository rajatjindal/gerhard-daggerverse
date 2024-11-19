// Displays system information
//
// Implements a few flavours: neofetch & fastfetch
// PR welcome with your favourite: https://beucismis.github.io/awesome-fetch/

package main

import (
	"context"
	"dagger/sysi/internal/dagger"
	"time"
)

type Sysi struct{}

// https://github.com/dylanaraps/neofetch `dagger call neofetch`
func (m *Sysi) Neofetch(ctx context.Context) (string, error) {
	return m.apk("neofetch").WithExec([]string{"neofetch"}).Stdout(ctx)
}

// https://github.com/fastfetch-cli/fastfetch `dagger call fastfetch`
func (m *Sysi) Fastfetch(ctx context.Context) (string, error) {
	return m.apk("fastfetch").WithExec([]string{"fastfetch", "--pipe", "false"}).Stdout(ctx)
}

func (m *Sysi) apk(pkg string) *dagger.Container {
	// https://hub.docker.com/_/alpine/tags
	return dag.Container().
		From("alpine:3.20.3@sha256:1e42bbe2508154c9126d48c2b8a75420c3544343bf86fd041fb7527e017a4b4a").
		WithExec([]string{"apk", "add", pkg}).
		WithEnvVariable("CACHE_BUSTED_AT", time.Now().String())
}
