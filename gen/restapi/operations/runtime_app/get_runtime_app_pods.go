// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppPodsHandlerFunc turns a function with the right signature into a get runtime app pods handler
type GetRuntimeAppPodsHandlerFunc func(GetRuntimeAppPodsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppPodsHandlerFunc) Handle(params GetRuntimeAppPodsParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppPodsHandler interface for that can handle valid get runtime app pods params
type GetRuntimeAppPodsHandler interface {
	Handle(GetRuntimeAppPodsParams) middleware.Responder
}

// NewGetRuntimeAppPods creates a new http.Handler for the get runtime app pods operation
func NewGetRuntimeAppPods(ctx *middleware.Context, handler GetRuntimeAppPodsHandler) *GetRuntimeAppPods {
	return &GetRuntimeAppPods{Context: ctx, Handler: handler}
}

/*GetRuntimeAppPods swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/pods runtimeApp getRuntimeAppPods

GetRuntimeAppPods get runtime app pods API

*/
type GetRuntimeAppPods struct {
	Context *middleware.Context
	Handler GetRuntimeAppPodsHandler
}

func (o *GetRuntimeAppPods) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppPodsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}