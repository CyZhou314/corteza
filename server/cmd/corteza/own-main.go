package main

import (
	"github.com/cyzhou314/corteza/server/app"
	"github.com/cyzhou314/corteza/server/pkg/cli"
	"github.com/cyzhou314/corteza/server/pkg/logger"
)

func main() {
	// Initialize logger before any other action
	logger.Init()

	cli.HandleError(app.New().Execute())
}
