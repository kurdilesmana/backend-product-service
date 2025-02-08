package main

import (
	"github.com/kurdilesmana/backend-product-service/cmd"
	"github.com/kurdilesmana/backend-product-service/deps"
)

func main() {
	dependency := deps.SetupDependencies()
	go func() {
		cmd.ExecuteGrpc(dependency)
	}()
	cmd.ExecuteHTTP(dependency)
}
