package app

// Registers all supported store backends
import (
	_ "github.com/cyzhou314/corteza/server/store/adapters/rdbms/drivers/mssql"
	_ "github.com/cyzhou314/corteza/server/store/adapters/rdbms/drivers/mysql"
	_ "github.com/cyzhou314/corteza/server/store/adapters/rdbms/drivers/postgres"
	_ "github.com/cyzhou314/corteza/server/store/adapters/rdbms/drivers/sqlite"
)
