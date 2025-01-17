package commands

import (
	"context"

	cs "github.com/cyzhou314/corteza/server/compose/service"
	"github.com/cyzhou314/corteza/server/federation/service"
	ss "github.com/cyzhou314/corteza/server/system/service"
	"github.com/spf13/cobra"
)

func commandSyncData(ctx context.Context) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, _ []string) {
		syncService := service.NewSync(
			&service.Syncer{},
			&service.Mapper{},
			service.DefaultSharedModule,
			cs.DefaultRecord,
			ss.DefaultUser,
			ss.DefaultRole)

		syncData := service.WorkerData(syncService, service.DefaultLogger)
		syncData.Watch(ctx, service.DefaultOptions.DataMonitorInterval, service.DefaultOptions.DataPageSize)
	}
}
