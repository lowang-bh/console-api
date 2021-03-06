// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppConfigMapHandlerFunc turns a function with the right signature into a get runtime app config map handler
type GetRuntimeAppConfigMapHandlerFunc func(GetRuntimeAppConfigMapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppConfigMapHandlerFunc) Handle(params GetRuntimeAppConfigMapParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppConfigMapHandler interface for that can handle valid get runtime app config map params
type GetRuntimeAppConfigMapHandler interface {
	Handle(GetRuntimeAppConfigMapParams) middleware.Responder
}

// NewGetRuntimeAppConfigMap creates a new http.Handler for the get runtime app config map operation
func NewGetRuntimeAppConfigMap(ctx *middleware.Context, handler GetRuntimeAppConfigMapHandler) *GetRuntimeAppConfigMap {
	return &GetRuntimeAppConfigMap{Context: ctx, Handler: handler}
}

/*GetRuntimeAppConfigMap swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/configmap runtimeApp getRuntimeAppConfigMap

GetRuntimeAppConfigMap get runtime app config map API

*/
type GetRuntimeAppConfigMap struct {
	Context *middleware.Context
	Handler GetRuntimeAppConfigMapHandler
}

func (o *GetRuntimeAppConfigMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppConfigMapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
