package glsp

import (
	contextpkg "context"
	"encoding/json"
)

type NotifyFunc func(method string, params any)
type CallFunc func(method string, params any, result any)
type ConnectionDetails struct {
	Id int
}

type Context struct {
	Method            string
	Params            json.RawMessage
	Notify            NotifyFunc
	Call              CallFunc
	Context           contextpkg.Context // can be nil
	ConnectionDetails ConnectionDetails
}

type Handler interface {
	Handle(context *Context) (result any, validMethod bool, validParams bool, err error)
}
