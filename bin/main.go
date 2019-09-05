package main

import (
	"github.com/kyamato/docker-machine-driver-generic"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(generic.NewDriver("", ""))
}
