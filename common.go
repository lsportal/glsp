package glsp

import (
	contextpkg "context"
	"encoding/json"
)

type NotifyFunc func(method string, params any)
type CallFunc func(method string, params any, result any)

type NotifyOtherFunc func(method string, params any, client string)
type CallOtherFunc func(method string, params any, result any, client string)
type ConnectionDetails struct {
	Id string
}

type Context struct {
	Method            string
	Params            json.RawMessage
	Notify            NotifyFunc
	NotifyOther       NotifyOtherFunc
	Call              CallFunc
	CallOther         CallOtherFunc
	Context           contextpkg.Context // can be nil
	ConnectionDetails ConnectionDetails
}

type Handler interface {
	Handle(context *Context) (result any, validMethod bool, validParams bool, err error)
}
