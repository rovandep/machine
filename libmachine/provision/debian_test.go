package provision

import (
	"testing"

	"github.com/code-ready/machine/drivers/fakedriver"
	"github.com/code-ready/machine/libmachine/auth"
	"github.com/code-ready/machine/libmachine/engine"
	"github.com/code-ready/machine/libmachine/provision/provisiontest"
	"github.com/code-ready/machine/libmachine/swarm"
)

func TestDebianDefaultStorageDriver(t *testing.T) {
	p := NewDebianProvisioner(&fakedriver.Driver{}).(*DebianProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay2" {
		t.Fatal("Default storage driver should be overlay2")
	}
}
