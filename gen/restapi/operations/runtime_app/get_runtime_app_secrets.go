// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppSecretsHandlerFunc turns a function with the right signature into a get runtime app secrets handler
type GetRuntimeAppSecretsHandlerFunc func(GetRuntimeAppSecretsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppSecretsHandlerFunc) Handle(params GetRuntimeAppSecretsParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppSecretsHandler interface for that can handle valid get runtime app secrets params
type GetRuntimeAppSecretsHandler interface {
	Handle(GetRuntimeAppSecretsParams) middleware.Responder
}

// NewGetRuntimeAppSecrets creates a new http.Handler for the get runtime app secrets operation
func NewGetRuntimeAppSecrets(ctx *middleware.Context, handler GetRuntimeAppSecretsHandler) *GetRuntimeAppSecrets {
	return &GetRuntimeAppSecrets{Context: ctx, Handler: handler}
}

/*GetRuntimeAppSecrets swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/secrets runtimeApp getRuntimeAppSecrets

GetRuntimeAppSecrets get runtime app secrets API

*/
type GetRuntimeAppSecrets struct {
	Context *middleware.Context
	Handler GetRuntimeAppSecretsHandler
}

func (o *GetRuntimeAppSecrets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppSecretsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}