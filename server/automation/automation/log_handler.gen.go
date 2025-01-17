package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// automation/automation/log_handler.yaml

import (
	"context"
	atypes "github.com/cyzhou314/corteza/server/automation/types"
	"github.com/cyzhou314/corteza/server/pkg/expr"
	"github.com/cyzhou314/corteza/server/pkg/wfexec"
)

var _ wfexec.ExecResponse

type (
	logHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h logHandler) register() {
	h.reg.AddFunctions(
		h.Debug(),
		h.Info(),
		h.Warn(),
		h.Error(),
	)
}

type (
	logDebugArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

// Debug function Log debug message
//
// expects implementation of debug function:
// func (h logHandler) debug(ctx context.Context, args *logDebugArgs) (err error) {
//    return
// }
func (h logHandler) Debug() *atypes.Function {
	return &atypes.Function{
		Ref:    "logDebug",
		Kind:   "function",
		Labels: map[string]string{"debug": "step", "logger": "step,workflow"},
		Meta: &atypes.FunctionMeta{
			Short: "Log debug message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &logDebugArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.debug(ctx, args)
		},
	}
}

type (
	logInfoArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

// Info function Log info message
//
// expects implementation of info function:
// func (h logHandler) info(ctx context.Context, args *logInfoArgs) (err error) {
//    return
// }
func (h logHandler) Info() *atypes.Function {
	return &atypes.Function{
		Ref:    "logInfo",
		Kind:   "function",
		Labels: map[string]string{"debug": "step", "logger": "step,workflow"},
		Meta: &atypes.FunctionMeta{
			Short: "Log info message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &logInfoArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.info(ctx, args)
		},
	}
}

type (
	logWarnArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

// Warn function Log warning message
//
// expects implementation of warn function:
// func (h logHandler) warn(ctx context.Context, args *logWarnArgs) (err error) {
//    return
// }
func (h logHandler) Warn() *atypes.Function {
	return &atypes.Function{
		Ref:    "logWarn",
		Kind:   "function",
		Labels: map[string]string{"debug": "step", "logger": "step,workflow"},
		Meta: &atypes.FunctionMeta{
			Short: "Log warning message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &logWarnArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.warn(ctx, args)
		},
	}
}

type (
	logErrorArgs struct {
		hasMessage bool
		Message    string

		hasFields bool
		Fields    map[string]string
	}
)

// Error function Log error message
//
// expects implementation of error function:
// func (h logHandler) error(ctx context.Context, args *logErrorArgs) (err error) {
//    return
// }
func (h logHandler) Error() *atypes.Function {
	return &atypes.Function{
		Ref:    "logError",
		Kind:   "function",
		Labels: map[string]string{"debug": "step", "logger": "step,workflow"},
		Meta: &atypes.FunctionMeta{
			Short: "Log error message",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "message",
				Types: []string{"String"}, Required: true,
			},
			{
				Name:  "fields",
				Types: []string{"KV"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &logErrorArgs{
					hasMessage: in.Has("message"),
					hasFields:  in.Has("fields"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			return out, h.error(ctx, args)
		},
	}
}
