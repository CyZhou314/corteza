package yaml

import (
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	setting struct {
		res *types.SettingValue
		ts  *resource.Timestamps
		us  *resource.Userstamps

		envoyConfig   *resource.EnvoyConfig
		encoderConfig *EncoderConfig
	}

	settingSet []*setting
)
