// Code generated by go-swagger; DO NOT EDIT.

package runtime_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppsHandlerFunc turns a function with the right signature into a get runtime apps handler
type GetRuntimeAppsHandlerFunc func(GetRuntimeAppsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppsHandlerFunc) Handle(params GetRuntimeAppsParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppsHandler interface for that can handle valid get runtime apps params
type GetRuntimeAppsHandler interface {
	Handle(GetRuntimeAppsParams) middleware.Responder
}

// NewGetRuntimeApps creates a new http.Handler for the get runtime apps operation
func NewGetRuntimeApps(ctx *middleware.Context, handler GetRuntimeAppsHandler) *GetRuntimeApps {
	return &GetRuntimeApps{Context: ctx, Handler: handler}
}

/*GetRuntimeApps swagger:route GET /api/v1/groups/{groupname}/runtime/apps runtimeApps getRuntimeApps

GetRuntimeApps get runtime apps API

*/
type GetRuntimeApps struct {
	Context *middleware.Context
	Handler GetRuntimeAppsHandler
}

func (o *GetRuntimeApps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
