package external

import (
	"os"
	"testing"

	"github.com/cyzhou314/corteza/server/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.SetDefault(logger.MakeDebugLogger())
	os.Exit(m.Run())
}
