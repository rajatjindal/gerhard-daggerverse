package main

import (
	"context"
	"dagger/dagrr/internal/dagger"
	"fmt"
)

type DagrrFly struct {
	// +private
	Dagrr *Dagrr
	// +private
	Flyio *dagger.Flyio
}

// App manifest: `dagger call on-flyio --token=env:FLY_API_TOKEN manifest file --path=fly.toml export --path=fly.toml`
func (m *DagrrFly) Manifest(
	// Disk size in GB
	//
	// +optional
	// +default="100GB"
	disk string,

	// VM size, see https://fly.io/docs/about/pricing/#compute
	//
	// +optional
	// +default="performance-2x"
	size string,
) *dagger.Directory {
	toml := fmt.Sprintf(`# https://fly.io/docs/reference/configuration/

app = "%s"

kill_signal = "SIGINT"
kill_timeout = 30

[build]
  image = "registry.dagger.io/engine:v%s"

[mounts]
  source = "dagger"
  destination = "/var/lib/dagger"
  initial_size = "%s"

[processes]
  dagger = "--addr unix:///var/run/buildkit/buildkitd.sock --addr tcp://0.0.0.0:2345"

[checks]
  [checks.http]
    grace_period = "3s"
    interval = "2s"
    port = 2345
    timeout = "1s"
    type = "tcp"

[[services]]
  internal_port = 2345
  protocol = "tcp"
  auto_stop_machines = false
  auto_start_machines = true
  min_machines_running = 1
  processes = ["dagger"]

  [[services.ports]]
    handlers = ["http"]
    port = 2345

[[vm]]
  size = "%s"
	`, m.Dagrr.App, m.Dagrr.Version, disk, size)

	return dag.Directory().WithNewFile("fly.toml", toml)
}

// Deploy with default manifest: `dagger call on-flyio --token=env:FLY_API_TOKEN deploy`
// Then: `export _EXPERIMENTAL_DAGGER_RUNNER_HOST=tcp://<APP_NAME>.internal:2345`
// Assumes https://fly.io/docs/networking/private-networking (clashes with Tailscale MagicDNS)
func (m *DagrrFly) Deploy(
	// +optional
	dir *dagger.Directory,
) (string, error) {
	ctx := context.Background()

	create, err := m.Flyio.Create(ctx, m.Dagrr.App)
	if err != nil {
		return create, err
	}

	if dir == nil {
		dir = m.Manifest("100GB", "performance-2x")
	}

	return m.Flyio.Deploy(ctx, dir)
}
