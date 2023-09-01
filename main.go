package main

import (
	"context"
	_ "embed"
)

//go:embed version
var Version string

func Run(ctx context.Context, target string, command string, ExpectedOutput string, username string, password string) (bool, string) {
	return true, "Successfully completed"
}
