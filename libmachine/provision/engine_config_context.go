package provision

import (
	"github.com/code-ready/machine/libmachine/auth"
	"github.com/code-ready/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.Options
	EngineOptions    engine.Options
	DockerOptionsDir string
}
