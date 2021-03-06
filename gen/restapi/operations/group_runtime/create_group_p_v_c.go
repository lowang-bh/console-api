// Code generated by go-swagger; DO NOT EDIT.

package group_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateGroupPVCHandlerFunc turns a function with the right signature into a create group p v c handler
type CreateGroupPVCHandlerFunc func(CreateGroupPVCParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateGroupPVCHandlerFunc) Handle(params CreateGroupPVCParams) middleware.Responder {
	return fn(params)
}

// CreateGroupPVCHandler interface for that can handle valid create group p v c params
type CreateGroupPVCHandler interface {
	Handle(CreateGroupPVCParams) middleware.Responder
}

// NewCreateGroupPVC creates a new http.Handler for the create group p v c operation
func NewCreateGroupPVC(ctx *middleware.Context, handler CreateGroupPVCHandler) *CreateGroupPVC {
	return &CreateGroupPVC{Context: ctx, Handler: handler}
}

/*CreateGroupPVC swagger:route POST /api/v1/groups/{groupname}/runtime/pvcs groupRuntime createGroupPVC

CreateGroupPVC create group p v c API

*/
type CreateGroupPVC struct {
	Context *middleware.Context
	Handler CreateGroupPVCHandler
}

func (o *CreateGroupPVC) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateGroupPVCParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
