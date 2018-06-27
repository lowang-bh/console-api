// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateRuntimeAppSecretHandlerFunc turns a function with the right signature into a update runtime app secret handler
type UpdateRuntimeAppSecretHandlerFunc func(UpdateRuntimeAppSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateRuntimeAppSecretHandlerFunc) Handle(params UpdateRuntimeAppSecretParams) middleware.Responder {
	return fn(params)
}

// UpdateRuntimeAppSecretHandler interface for that can handle valid update runtime app secret params
type UpdateRuntimeAppSecretHandler interface {
	Handle(UpdateRuntimeAppSecretParams) middleware.Responder
}

// NewUpdateRuntimeAppSecret creates a new http.Handler for the update runtime app secret operation
func NewUpdateRuntimeAppSecret(ctx *middleware.Context, handler UpdateRuntimeAppSecretHandler) *UpdateRuntimeAppSecret {
	return &UpdateRuntimeAppSecret{Context: ctx, Handler: handler}
}

/*UpdateRuntimeAppSecret swagger:route PUT /api/v1/groups/{groupname}/runtime/apps/{appname}/secrets/{secretname} runtimeApp updateRuntimeAppSecret

UpdateRuntimeAppSecret update runtime app secret API

*/
type UpdateRuntimeAppSecret struct {
	Context *middleware.Context
	Handler UpdateRuntimeAppSecretHandler
}

func (o *UpdateRuntimeAppSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateRuntimeAppSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
