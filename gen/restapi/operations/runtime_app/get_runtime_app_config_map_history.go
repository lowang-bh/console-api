// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppConfigMapHistoryHandlerFunc turns a function with the right signature into a get runtime app config map history handler
type GetRuntimeAppConfigMapHistoryHandlerFunc func(GetRuntimeAppConfigMapHistoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppConfigMapHistoryHandlerFunc) Handle(params GetRuntimeAppConfigMapHistoryParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppConfigMapHistoryHandler interface for that can handle valid get runtime app config map history params
type GetRuntimeAppConfigMapHistoryHandler interface {
	Handle(GetRuntimeAppConfigMapHistoryParams) middleware.Responder
}

// NewGetRuntimeAppConfigMapHistory creates a new http.Handler for the get runtime app config map history operation
func NewGetRuntimeAppConfigMapHistory(ctx *middleware.Context, handler GetRuntimeAppConfigMapHistoryHandler) *GetRuntimeAppConfigMapHistory {
	return &GetRuntimeAppConfigMapHistory{Context: ctx, Handler: handler}
}

/*GetRuntimeAppConfigMapHistory swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/configmap/history runtimeApp getRuntimeAppConfigMapHistory

GetRuntimeAppConfigMapHistory get runtime app config map history API

*/
type GetRuntimeAppConfigMapHistory struct {
	Context *middleware.Context
	Handler GetRuntimeAppConfigMapHistoryHandler
}

func (o *GetRuntimeAppConfigMapHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppConfigMapHistoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
