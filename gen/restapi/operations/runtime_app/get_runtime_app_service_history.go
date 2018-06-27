// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRuntimeAppServiceHistoryHandlerFunc turns a function with the right signature into a get runtime app service history handler
type GetRuntimeAppServiceHistoryHandlerFunc func(GetRuntimeAppServiceHistoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeAppServiceHistoryHandlerFunc) Handle(params GetRuntimeAppServiceHistoryParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeAppServiceHistoryHandler interface for that can handle valid get runtime app service history params
type GetRuntimeAppServiceHistoryHandler interface {
	Handle(GetRuntimeAppServiceHistoryParams) middleware.Responder
}

// NewGetRuntimeAppServiceHistory creates a new http.Handler for the get runtime app service history operation
func NewGetRuntimeAppServiceHistory(ctx *middleware.Context, handler GetRuntimeAppServiceHistoryHandler) *GetRuntimeAppServiceHistory {
	return &GetRuntimeAppServiceHistory{Context: ctx, Handler: handler}
}

/*GetRuntimeAppServiceHistory swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/service/history runtimeApp getRuntimeAppServiceHistory

GetRuntimeAppServiceHistory get runtime app service history API

*/
type GetRuntimeAppServiceHistory struct {
	Context *middleware.Context
	Handler GetRuntimeAppServiceHistoryHandler
}

func (o *GetRuntimeAppServiceHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuntimeAppServiceHistoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
