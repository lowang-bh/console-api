// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateOrUpdateRuntimeAppConfigMapHandlerFunc turns a function with the right signature into a create or update runtime app config map handler
type CreateOrUpdateRuntimeAppConfigMapHandlerFunc func(CreateOrUpdateRuntimeAppConfigMapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrUpdateRuntimeAppConfigMapHandlerFunc) Handle(params CreateOrUpdateRuntimeAppConfigMapParams) middleware.Responder {
	return fn(params)
}

// CreateOrUpdateRuntimeAppConfigMapHandler interface for that can handle valid create or update runtime app config map params
type CreateOrUpdateRuntimeAppConfigMapHandler interface {
	Handle(CreateOrUpdateRuntimeAppConfigMapParams) middleware.Responder
}

// NewCreateOrUpdateRuntimeAppConfigMap creates a new http.Handler for the create or update runtime app config map operation
func NewCreateOrUpdateRuntimeAppConfigMap(ctx *middleware.Context, handler CreateOrUpdateRuntimeAppConfigMapHandler) *CreateOrUpdateRuntimeAppConfigMap {
	return &CreateOrUpdateRuntimeAppConfigMap{Context: ctx, Handler: handler}
}

/*CreateOrUpdateRuntimeAppConfigMap swagger:route PUT /api/v1/groups/{groupname}/runtime/apps/{appname}/configmap runtimeApp createOrUpdateRuntimeAppConfigMap

CreateOrUpdateRuntimeAppConfigMap create or update runtime app config map API

*/
type CreateOrUpdateRuntimeAppConfigMap struct {
	Context *middleware.Context
	Handler CreateOrUpdateRuntimeAppConfigMapHandler
}

func (o *CreateOrUpdateRuntimeAppConfigMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateOrUpdateRuntimeAppConfigMapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}