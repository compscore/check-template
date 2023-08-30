package main

import (
	"context"
	_ "embed"

	"github.com/compscore/compscore/pkg/structs"
)

//go:embed version
var Version string

func Run(ctx context.Context, target string, command string, expectedOutput string, credentials structs.Credentials_s) (bool, string) {
	return true, "Successfully completed"
}
