//go:generate go run ./generate.go

package main

import (
	"github.com/LuxChanLu/pulumi-gomigrate/pkg/provider"
	"github.com/LuxChanLu/pulumi-gomigrate/pkg/version"
)

var providerName = "gomigrate"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema)
}
