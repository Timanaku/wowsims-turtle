package main

import (
	"github.com/isfir/wowsims-turtle/cmd/wowsimcli/cmd"
	"github.com/isfir/wowsims-turtle/sim"
)

func init() {
	sim.RegisterAll()
}

// Version information.
// This variable is set by the makefile in the release process.
var Version string

func main() {
	cmd.Execute(Version)
}
