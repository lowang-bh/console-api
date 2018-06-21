// Code generated by go-swagger; DO NOT EDIT.

package runtime_app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateRuntimeAppSecretHandlerFunc turns a function with the right signature into a create runtime app secret handler
type CreateRuntimeAppSecretHandlerFunc func(CreateRuntimeAppSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRuntimeAppSecretHandlerFunc) Handle(params CreateRuntimeAppSecretParams) middleware.Responder {
	return fn(params)
}

// CreateRuntimeAppSecretHandler interface for that can handle valid create runtime app secret params
type CreateRuntimeAppSecretHandler interface {
	Handle(CreateRuntimeAppSecretParams) middleware.Responder
}

// NewCreateRuntimeAppSecret creates a new http.Handler for the create runtime app secret operation
func NewCreateRuntimeAppSecret(ctx *middleware.Context, handler CreateRuntimeAppSecretHandler) *CreateRuntimeAppSecret {
	return &CreateRuntimeAppSecret{Context: ctx, Handler: handler}
}

/*CreateRuntimeAppSecret swagger:route POST /api/v1/groups/{groupname}/runtime/apps/{appname}/secrets runtimeAppSecrets createRuntimeAppSecret

CreateRuntimeAppSecret create runtime app secret API

*/
type CreateRuntimeAppSecret struct {
	Context *middleware.Context
	Handler CreateRuntimeAppSecretHandler
}

func (o *CreateRuntimeAppSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRuntimeAppSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
