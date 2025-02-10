package main

import (
	"github.com/kurdilesmana/backend-product-service/cmd"
	"github.com/kurdilesmana/backend-product-service/deps"
)

func main() {
	dependency := deps.SetupDependencies()
	cmd.ExecuteHTTP(dependency)
}
