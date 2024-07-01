// Displays system information
//
// Implements a few flavours: neofetch & fastfetch
// PR welcome with your favourite: https://beucismis.github.io/awesome-fetch/

package main

import (
	"context"
	"time"
)

type Sysi struct{}

// https://github.com/dylanaraps/neofetch
func (m *Sysi) Neofetch(ctx context.Context) (string, error) {
	return m.apk("neofetch").WithExec([]string{"neofetch"}).Stdout(ctx)
}

// https://github.com/fastfetch-cli/fastfetch
func (m *Sysi) Fastfetch(ctx context.Context) (string, error) {
	return m.apk("fastfetch").WithExec([]string{"fastfetch", "--pipe", "false"}).Stdout(ctx)
}

func (m *Sysi) apk(pkg string) *Container {
	return dag.Container().
		From("alpine:20240606@sha256:166710df254975d4a6c4c407c315951c22753dcaa829e020a3fd5d18fff70dd2").
		WithExec([]string{"apk", "add", pkg}).
		WithEnvVariable("CACHE_BUSTED_AT", time.Now().String())
}
