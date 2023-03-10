package commands

import (
	"context"

	"github.com/cyzhou314/corteza/server/pkg/cli"
	"github.com/spf13/cobra"
)

type (
	serviceInitializer interface {
		InitServices(ctx context.Context) error
	}
)

func commandPreRunInitService(app serviceInitializer) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, _ []string) error {
		return app.InitServices(cli.Context())
	}
}
