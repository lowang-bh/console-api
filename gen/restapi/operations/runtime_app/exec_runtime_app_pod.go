// Code generated by go-swagger; DO NOT EDIT.

package runtime_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ExecRuntimeAppPodHandlerFunc turns a function with the right signature into a exec runtime app pod handler
type ExecRuntimeAppPodHandlerFunc func(ExecRuntimeAppPodParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ExecRuntimeAppPodHandlerFunc) Handle(params ExecRuntimeAppPodParams) middleware.Responder {
	return fn(params)
}

// ExecRuntimeAppPodHandler interface for that can handle valid exec runtime app pod params
type ExecRuntimeAppPodHandler interface {
	Handle(ExecRuntimeAppPodParams) middleware.Responder
}

// NewExecRuntimeAppPod creates a new http.Handler for the exec runtime app pod operation
func NewExecRuntimeAppPod(ctx *middleware.Context, handler ExecRuntimeAppPodHandler) *ExecRuntimeAppPod {
	return &ExecRuntimeAppPod{Context: ctx, Handler: handler}
}

/*ExecRuntimeAppPod swagger:route GET /api/v1/groups/{groupname}/runtime/apps/{appname}/pods/{podname}/exec runtimeApp execRuntimeAppPod

长连接

*/
type ExecRuntimeAppPod struct {
	Context *middleware.Context
	Handler ExecRuntimeAppPodHandler
}

func (o *ExecRuntimeAppPod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExecRuntimeAppPodParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}