// Code generated by go-swagger; DO NOT EDIT.

package runtime_apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppHandlerFunc turns a function with the right signature into a get runtime app handler
type GetRuntimeAppHandlerFunc func(GetRuntimeAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppHandlerFunc) Handle(params GetRuntimeAppParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppHandler interface for that can handle valid get runtime app params
type GetRuntimeAppHandler interface {
	Handle(GetRuntimeAppParams) middleware.Responder
}

// NewGetRuntimeApp creates a new http.Handler for the get runtime app operation
func NewGetRuntimeApp(ctx *middleware.Context, handler GetRuntimeAppHandler) *GetRuntimeApp {
	return &GetRuntimeApp{Context: ctx, Handler: handler}
}

/*GetRuntimeApp swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname} runtimeApps getRuntimeApp

GetRuntimeApp get runtime app API

*/
type GetRuntimeApp struct {
	Context *middleware.Context
	Handler GetRuntimeAppHandler
}

func (o *GetRuntimeApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
