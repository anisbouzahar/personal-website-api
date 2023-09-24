//go:build tools

package main

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/wire"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/spf13/cobra"
	_ "github.com/swaggo/swag"
	_ "github.com/tebeka/go2xunit"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "golang.org/x/tools/cmd/stringer"
)
