package types

import (
	"time"

	"github.com/cyzhou314/corteza/server/pkg/filter"
)

type (
	SharedModule struct {
		ID                         uint64         `json:"moduleID,string"`
		NodeID                     uint64         `json:"nodeID,string"`
		Handle                     string         `json:"handle"`
		Name                       string         `json:"name"`
		ExternalFederationModuleID uint64         `json:"externalFederationModuleID,string"`
		Fields                     ModuleFieldSet `json:"fields"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		CreatedBy uint64     `json:"createdBy,string" `
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		UpdatedBy uint64     `json:"updatedBy,string,omitempty" `
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
		DeletedBy uint64     `json:"deletedBy,string,omitempty" `
	}

	SharedModuleFilter struct {
		NodeID                     uint64 `json:"nodeID,string"`
		ExternalFederationModuleID uint64 `json:"externalFederationModuleID,string"`

		Handle string `json:"handle"`
		Name   string `json:"name"`
		Query  string `json:"query"`

		Check func(*SharedModule) (bool, error) `json:"-"`

		filter.Sorting
		filter.Paging
	}
)
