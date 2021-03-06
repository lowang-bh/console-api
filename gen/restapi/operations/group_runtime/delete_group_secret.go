// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteGroupSecretHandlerFunc turns a function with the right signature into a delete group secret handler
type DeleteGroupSecretHandlerFunc func(DeleteGroupSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteGroupSecretHandlerFunc) Handle(params DeleteGroupSecretParams) middleware.Responder {
	return fn(params)
}

// DeleteGroupSecretHandler interface for that can handle valid delete group secret params
type DeleteGroupSecretHandler interface {
	Handle(DeleteGroupSecretParams) middleware.Responder
}

// NewDeleteGroupSecret creates a new http.Handler for the delete group secret operation
func NewDeleteGroupSecret(ctx *middleware.Context, handler DeleteGroupSecretHandler) *DeleteGroupSecret {
	return &DeleteGroupSecret{Context: ctx, Handler: handler}
}

/*DeleteGroupSecret swagger:route DELETE /api/v1/groups/{groupname}/runtime/secrets/{secretname} groupRuntime deleteGroupSecret

DeleteGroupSecret delete group secret API

*/
type DeleteGroupSecret struct {
	Context *middleware.Context
	Handler DeleteGroupSecretHandler
}

func (o *DeleteGroupSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteGroupSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
