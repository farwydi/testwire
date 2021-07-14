// +build wireinject

package main

import (
	"github.com/google/wire"
	"testwire/rep"
)

func setup() app {
	panic(wire.Build(rep.NewRep, newApp))
}
