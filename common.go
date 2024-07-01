package glsp

import (
	contextpkg "context"
	"encoding/json"
)

type (
	NotifyFunc func(method string, params any)
	CallFunc   func(method string, params any, result any)
)

type Context struct {
	Method      string
	Params      json.RawMessage
	Notify      NotifyFunc
	Call        CallFunc
	CallOther   CallFunc
	NotifyOther NotifyFunc
	Context     contextpkg.Context // can be nil
}

type Handler interface {
	Handle(context *Context) (result any, validMethod bool, validParams bool, err error)
}
