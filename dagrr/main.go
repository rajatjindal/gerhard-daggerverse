// Manages Dagger on a bunch of platforms

package main

import (
	"dagger/dagrr/internal/dagger"
	"strings"
	"time"
)

type Dagrr struct {
	// +private
	Version string
	// +private
	App string
}

func New(
	// Dagger version to use: `--version=0.11.9`
	//
	// +optional
	// https://github.com/dagger/dagger/blob/main/CHANGELOG.md
	// +default="0.12.5"
	version string,

	// App name, defaults to version & date: `--app=dagger-v0119-2024-07-03`
	//
	// +optional
	app string,
) *Dagrr {

	m := &Dagrr{
		Version: version,
	}

	if app == "" {
		app = strings.Join([]string{
			"dagger",
			m.versionUrlized(),
			time.Now().Format("2006-01-02"),
		}, "-")
	}
	m.App = app

	return m
}

// Manages Dagger on Fly.io
func (m *Dagrr) OnFlyio(
	// fly auth token: `--token=env:FLY_API_TOKEN`
	token *dagger.Secret,

	// Fly.io org name
	//
	// +optional
	// +default="personal"
	org string,
) *DagrrFly {
	return &DagrrFly{
		Dagrr: m,
		Flyio: dag.Flyio(token, dagger.FlyioOpts{
			Org: org,
		}),
	}
}

func (m *Dagrr) versionUrlized() string {
	return "v" + strings.ReplaceAll(m.Version, ".", "")
}
