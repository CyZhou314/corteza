package resource

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

// Definitions file that controls how this file is generated:
// - automation.workflow.yaml

import (
	"github.com/cyzhou314/corteza/server/automation/types"
)

// AutomationWorkflowResourceTranslationReferences generates Locale references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func AutomationWorkflowResourceTranslationReferences(workflow string) (res *Ref, pp []*Ref, err error) {
	res = &Ref{ResourceType: types.WorkflowResourceType, Identifiers: MakeIdentifiers(workflow)}

	return
}
