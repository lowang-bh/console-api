// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGroupSecretHandlerFunc turns a function with the right signature into a get group secret handler
type GetGroupSecretHandlerFunc func(GetGroupSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupSecretHandlerFunc) Handle(params GetGroupSecretParams) middleware.Responder {
	return fn(params)
}

// GetGroupSecretHandler interface for that can handle valid get group secret params
type GetGroupSecretHandler interface {
	Handle(GetGroupSecretParams) middleware.Responder
}

// NewGetGroupSecret creates a new http.Handler for the get group secret operation
func NewGetGroupSecret(ctx *middleware.Context, handler GetGroupSecretHandler) *GetGroupSecret {
	return &GetGroupSecret{Context: ctx, Handler: handler}
}

/*GetGroupSecret swagger:route GET /api/v1/groups/{groupname}/runtime/secrets/{secretname} groupRuntime getGroupSecret

GetGroupSecret get group secret API

*/
type GetGroupSecret struct {
	Context *middleware.Context
	Handler GetGroupSecretHandler
}

func (o *GetGroupSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGroupSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
