package main

import (
	"fmt"
	"keyhold/internal/args"
	"keyhold/internal/config"
	"keyhold/internal/logging"
)

func main() {
	args.Parse()
	logging.Init(args.IsProduction())
	err := config.Init()
	if err != nil {
		logging.Logger.Fatal(fmt.Errorf("initializing config: %w", err))
	}

	println("hello world")
}
