package sqlite

import (
	"context"

	"github.com/cyzhou314/corteza/server/pkg/logger"
	"github.com/cyzhou314/corteza/server/store/adapters/rdbms"
	"github.com/cyzhou314/corteza/server/store/adapters/rdbms/drivers/sqlite"
	"github.com/jmoiron/sqlx"
)

func Setup(ctx context.Context, dsn string) (_ *sqlx.DB, err error) {
	var (
		cfg *rdbms.ConnConfig
	)

	cfg, err = sqlite.NewConfig(dsn)
	if err != nil {
		return
	}

	return rdbms.Connect(ctx, logger.MakeDebugLogger(), cfg)
}
