package main

import (
	"keyhold/internal/args"
	"keyhold/internal/logging"
)

func main() {
	args.Parse()
	logging.Init(args.IsProduction())

	println("hello world")
}
