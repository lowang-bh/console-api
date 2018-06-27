package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/laincloud/console-api/gen/restapi/operations/ping"
)

// Ping show the health status of the server
func Ping(params ping.PingParams) middleware.Responder {
	return ping.NewPingOK().WithPayload("OK")
}
