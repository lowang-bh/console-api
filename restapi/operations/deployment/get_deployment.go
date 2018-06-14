// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDeploymentHandlerFunc turns a function with the right signature into a get deployment handler
type GetDeploymentHandlerFunc func(GetDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeploymentHandlerFunc) Handle(params GetDeploymentParams) middleware.Responder {
	return fn(params)
}

// GetDeploymentHandler interface for that can handle valid get deployment params
type GetDeploymentHandler interface {
	Handle(GetDeploymentParams) middleware.Responder
}

// NewGetDeployment creates a new http.Handler for the get deployment operation
func NewGetDeployment(ctx *middleware.Context, handler GetDeploymentHandler) *GetDeployment {
	return &GetDeployment{Context: ctx, Handler: handler}
}

/*GetDeployment swagger:route GET /groups/{group}/apps/{app}/deployment deployment getDeployment

GetDeployment get deployment API

*/
type GetDeployment struct {
	Context *middleware.Context
	Handler GetDeploymentHandler
}

func (o *GetDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}