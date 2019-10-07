package main

import (
	"github.com/kyamato/docker-machine-driver-dummy"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(dummy.NewDriver("", ""))
}
