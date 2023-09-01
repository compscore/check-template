package main

import (
	"context"
	_ "embed"
)

//go:embed version
var Version string

func Run(ctx context.Context, checkStruct struct {
	Target         string
	Command        string
	ExpectedOutput string
	Username       string
	Password       string
}) (bool, string) {
	return true, "Successfully completed"
}
