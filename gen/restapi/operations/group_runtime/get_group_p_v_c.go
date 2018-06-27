// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGroupPVCHandlerFunc turns a function with the right signature into a get group p v c handler
type GetGroupPVCHandlerFunc func(GetGroupPVCParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupPVCHandlerFunc) Handle(params GetGroupPVCParams) middleware.Responder {
	return fn(params)
}

// GetGroupPVCHandler interface for that can handle valid get group p v c params
type GetGroupPVCHandler interface {
	Handle(GetGroupPVCParams) middleware.Responder
}

// NewGetGroupPVC creates a new http.Handler for the get group p v c operation
func NewGetGroupPVC(ctx *middleware.Context, handler GetGroupPVCHandler) *GetGroupPVC {
	return &GetGroupPVC{Context: ctx, Handler: handler}
}

/*GetGroupPVC swagger:route GET /api/v1/groups/{groupname}/runtime/pvcs/{pvcname} groupRuntime getGroupPVC

GetGroupPVC get group p v c API

*/
type GetGroupPVC struct {
	Context *middleware.Context
	Handler GetGroupPVCHandler
}

func (o *GetGroupPVC) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGroupPVCParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
