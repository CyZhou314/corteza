package store

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	setting struct {
		cfg *EncoderConfig

		res *resource.Setting
		st  *types.SettingValue

		ux *userIndex
	}
)
