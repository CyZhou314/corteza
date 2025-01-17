package store

import (
	"fmt"

	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"github.com/cyzhou314/corteza/server/pkg/rbac"
	"github.com/cyzhou314/corteza/server/system/types"
)

type (
	rbacRule struct {
		cfg *EncoderConfig

		rule *rbac.Rule

		// point to the rbac rule
		refRbacResource string
		refRbacRes      *resource.Ref

		refPathRes []*resource.Ref

		refRole *resource.Ref
		role    *types.Role
	}
)

func rbacResourceErrUnidentifiable(ii resource.Identifiers) error {
	return fmt.Errorf("rbac resource unidentifiable %v", ii.StringSlice())
}
