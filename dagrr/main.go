// Manages Dagger Engines on a bunch of platforms

package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	// https://github.com/dagger/dagger/blob/main/CHANGELOG.md
	latestVersion = "0.11.9"
)

type Dagrr struct {
	// +private
	Version string
}

func New(
	// Dagger version to use: `--version=0.11.9`
	//
	// +optional
	version string,
) *Dagrr {
	if version == "" {
		version = latestVersion
	}

	return &Dagrr{
		Version: version,
	}
}

// Deploys Dagger on Fly.io
func (m *Dagrr) Flyio(
	// fly auth token: `--token=env:FLY_API_TOKEN`
	token *Secret,

	// App name, defaults to version & date: `--app=dagger-v0119-2024-07-03`
	//
	// +optional
	app string,
) (string, error) {
	if app == "" {
		app = strings.Join([]string{
			"dagger",
			m.versionUrlized(),
			time.Now().Format("2006-01-02"),
		}, "-")
	}

	ctx := context.Background()

	create, err := dag.Flyio(token).Create(ctx, app)
	if err != nil {
		return create, err
	}

	mount := strings.ReplaceAll(app, "-", "_")
	toml := fmt.Sprintf(`# https://fly.io/docs/reference/configuration/

app = "%s"
primary_region = "ams"

kill_signal = "SIGINT"
kill_timeout = 30

[build]
  image = "registry.dagger.io/engine:v%s"

[mounts]
  source = "%s"
  destination = "/var/lib/dagger"
  initial_size = "100GB"

[processes]
  dagger = "--addr unix:///var/run/buildkit/buildkitd.sock --addr tcp://0.0.0.0:1750"

[checks]
  [checks.http]
    grace_period = "3s"
    interval = "2s"
    port = 1750
    timeout = "1s"
    type = "tcp"

[[services]]
  internal_port = 1750
  protocol = "tcp"
  auto_stop_machines = true
  auto_start_machines = true
  min_machines_running = 1
  processes = ["dagger"]

  [[services.ports]]
    handlers = ["http"]
    port = 1750

[[vm]]
  size = "performance-2x"
	`, app, m.Version, mount)
	// return toml, nil

	return dag.Flyio(token).Deploy(ctx, dag.Directory().WithNewFile("fly.toml", toml))
}

func (m *Dagrr) versionUrlized() string {
	return "v" + strings.ReplaceAll(m.Version, ".", "")
}
