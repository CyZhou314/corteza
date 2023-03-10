package decoder

import (
	"github.com/cyzhou314/corteza/server/compose/types"
)

type (
	ComposeRecord struct {
		types.Record
	}
	ComposeRecordSet []*ComposeRecord

	ComposeRecordFilter struct {
		types.RecordFilter
	}
)
