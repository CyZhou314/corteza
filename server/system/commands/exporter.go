package commands

import (
	"context"
	"io"
	"os"
	"path"
	"strings"

	"github.com/cyzhou314/corteza/server/compose/service"
	"github.com/cyzhou314/corteza/server/pkg/auth"
	"github.com/cyzhou314/corteza/server/pkg/dal"
	"github.com/cyzhou314/corteza/server/pkg/envoy/yaml"
	"github.com/spf13/cobra"

	"github.com/cyzhou314/corteza/server/pkg/cli"
	"github.com/cyzhou314/corteza/server/pkg/envoy"
	su "github.com/cyzhou314/corteza/server/pkg/envoy/store"
	"github.com/cyzhou314/corteza/server/store"
)

func Export(ctx context.Context, storeInit func(ctx context.Context) (store.Storer, error)) *cobra.Command {
	var (
		output string
	)

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export",
		Long:  `Export data to YAML files.`,

		Run: func(cmd *cobra.Command, args []string) {
			ctx = auth.SetIdentityToContext(ctx, auth.ServiceUser())

			var (
				f = su.NewDecodeFilter()
			)

			s, err := storeInit(ctx)
			cli.HandleError(err)

			// init dal models
			err = service.DalModelReload(ctx, s, dal.Service())
			cli.HandleError(err)

			// TODO: init rbac, which is required by compose reocrd export

			f = f.FromResource(args...)

			sd := su.Decoder()
			nn, err := sd.Decode(ctx, s, dal.Service(), f)
			cli.HandleError(err)

			ye := yaml.NewYamlEncoder(&yaml.EncoderConfig{
				MappedOutput: true,
				// CompactOutput: true,
			})
			bld := envoy.NewBuilder(ye)
			g, err := bld.Build(ctx, nn...)
			cli.HandleError(err)
			err = envoy.Encode(ctx, g, ye)
			cli.HandleError(err)
			ss := ye.Stream()
			_ = ss

			makeFN := func(base, res string) string {
				pp := strings.Split(strings.Trim(res, ":"), ":")
				name := strings.Join(pp, "_") + ".yaml"
				return path.Join(base, name)
			}

			for _, s := range ss {
				f, err := os.Create(makeFN(output, s.Resource))
				cli.HandleError(err)
				defer f.Close()

				io.Copy(f, s.Source)
			}
		},
	}

	cmd.Flags().StringVarP(&output, "out", "o", "./", "The directory to write output files to")

	return cmd
}
