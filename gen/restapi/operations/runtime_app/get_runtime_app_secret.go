// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppSecretHandlerFunc turns a function with the right signature into a get runtime app secret handler
type GetRuntimeAppSecretHandlerFunc func(GetRuntimeAppSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppSecretHandlerFunc) Handle(params GetRuntimeAppSecretParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppSecretHandler interface for that can handle valid get runtime app secret params
type GetRuntimeAppSecretHandler interface {
	Handle(GetRuntimeAppSecretParams) middleware.Responder
}

// NewGetRuntimeAppSecret creates a new http.Handler for the get runtime app secret operation
func NewGetRuntimeAppSecret(ctx *middleware.Context, handler GetRuntimeAppSecretHandler) *GetRuntimeAppSecret {
	return &GetRuntimeAppSecret{Context: ctx, Handler: handler}
}

/*GetRuntimeAppSecret swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/secrets/{secretname} runtimeApp getRuntimeAppSecret

GetRuntimeAppSecret get runtime app secret API

*/
type GetRuntimeAppSecret struct {
	Context *middleware.Context
	Handler GetRuntimeAppSecretHandler
}

func (o *GetRuntimeAppSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
