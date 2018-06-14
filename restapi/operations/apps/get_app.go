// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAppHandlerFunc turns a function with the right signature into a get app handler
type GetAppHandlerFunc func(GetAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppHandlerFunc) Handle(params GetAppParams) middleware.Responder {
	return fn(params)
}

// GetAppHandler interface for that can handle valid get app params
type GetAppHandler interface {
	Handle(GetAppParams) middleware.Responder
}

// NewGetApp creates a new http.Handler for the get app operation
func NewGetApp(ctx *middleware.Context, handler GetAppHandler) *GetApp {
	return &GetApp{Context: ctx, Handler: handler}
}

/*GetApp swagger:route GET /groups/{group}/apps/{app} apps getApp

GetApp get app API

*/
type GetApp struct {
	Context *middleware.Context
	Handler GetAppHandler
}

func (o *GetApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
