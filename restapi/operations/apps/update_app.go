// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateAppHandlerFunc turns a function with the right signature into a update app handler
type UpdateAppHandlerFunc func(UpdateAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateAppHandlerFunc) Handle(params UpdateAppParams) middleware.Responder {
	return fn(params)
}

// UpdateAppHandler interface for that can handle valid update app params
type UpdateAppHandler interface {
	Handle(UpdateAppParams) middleware.Responder
}

// NewUpdateApp creates a new http.Handler for the update app operation
func NewUpdateApp(ctx *middleware.Context, handler UpdateAppHandler) *UpdateApp {
	return &UpdateApp{Context: ctx, Handler: handler}
}

/*UpdateApp swagger:route PUT /groups/{group}/apps/{app} apps updateApp

UpdateApp update app API

*/
type UpdateApp struct {
	Context *middleware.Context
	Handler UpdateAppHandler
}

func (o *UpdateApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
