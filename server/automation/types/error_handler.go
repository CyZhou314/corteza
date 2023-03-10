package types

import (
	"context"
	"github.com/cyzhou314/corteza/server/pkg/wfexec"
)

type (
	errorHandlerStep struct {
		wfexec.StepIdentifier
		handler wfexec.Step
	}
)

func ErrorHandlerStep(h wfexec.Step) *errorHandlerStep {
	return &errorHandlerStep{handler: h}
}

// Executes prompt step
func (h errorHandlerStep) Exec(_ context.Context, _ *wfexec.ExecRequest) (wfexec.ExecResponse, error) {
	return wfexec.ErrorHandler(h.handler), nil
}
