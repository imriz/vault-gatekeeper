package main

import (
	"runtime"

	// import schedulers before cmd
	_ "github.com/imriz/vault-gatekeeper/scheduler/ecs"
	_ "github.com/imriz/vault-gatekeeper/scheduler/mesos"

	"github.com/imriz/vault-gatekeeper/cmd"
)

var (
	Version   = "--dev--"
	BuildTime = "--dev--"
)

func main() {
	cmd.Version = Version
	cmd.RootCmd.Version = Version
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
