package main

import (
	"github.com/kyamato/generic"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(generic.NewDriver("", ""))
}
