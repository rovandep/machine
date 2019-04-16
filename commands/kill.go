package commands

import "github.com/code-ready/machine/libmachine"

func cmdKill(c CommandLine, api libmachine.API) error {
	return runAction("kill", c, api)
}
