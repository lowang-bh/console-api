// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/laincloud/console-api/gen/restapi/operations"
	"github.com/laincloud/console-api/gen/restapi/operations/build_app"
	"github.com/laincloud/console-api/gen/restapi/operations/build_apps"
	"github.com/laincloud/console-api/gen/restapi/operations/cluster"
	"github.com/laincloud/console-api/gen/restapi/operations/group_runtime"
	"github.com/laincloud/console-api/gen/restapi/operations/groups"
	"github.com/laincloud/console-api/gen/restapi/operations/ping"
	"github.com/laincloud/console-api/gen/restapi/operations/runtime_app"
	"github.com/laincloud/console-api/gen/restapi/operations/runtime_apps"
	"github.com/laincloud/console-api/handlers"
)

//go:generate swagger generate server --target ../gen --name  --spec ../swagger.json

func configureFlags(api *operations.ConsoleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ConsoleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TxtProducer = runtime.TextProducer()

	api.BuildAppsCreateBuildAppHandler = build_apps.CreateBuildAppHandlerFunc(func(params build_apps.CreateBuildAppParams) middleware.Responder {
		return middleware.NotImplemented("operation build_apps.CreateBuildApp has not yet been implemented")
	})
	api.GroupsCreateGroupHandler = groups.CreateGroupHandlerFunc(func(params groups.CreateGroupParams) middleware.Responder {
		return middleware.NotImplemented("operation groups.CreateGroup has not yet been implemented")
	})
	api.GroupRuntimeCreateGroupPVCHandler = group_runtime.CreateGroupPVCHandlerFunc(func(params group_runtime.CreateGroupPVCParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.CreateGroupPVC has not yet been implemented")
	})
	api.GroupRuntimeCreateGroupSecretHandler = group_runtime.CreateGroupSecretHandlerFunc(func(params group_runtime.CreateGroupSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.CreateGroupSecret has not yet been implemented")
	})
	api.GroupRuntimeCreateOrUpdateGroupIngressHandler = group_runtime.CreateOrUpdateGroupIngressHandlerFunc(func(params group_runtime.CreateOrUpdateGroupIngressParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.CreateOrUpdateGroupIngress has not yet been implemented")
	})
	api.RuntimeAppCreateOrUpdateRuntimeAppConfigMapHandler = runtime_app.CreateOrUpdateRuntimeAppConfigMapHandlerFunc(func(params runtime_app.CreateOrUpdateRuntimeAppConfigMapParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.CreateOrUpdateRuntimeAppConfigMap has not yet been implemented")
	})
	api.RuntimeAppCreateOrUpdateRuntimeAppServiceHandler = runtime_app.CreateOrUpdateRuntimeAppServiceHandlerFunc(func(params runtime_app.CreateOrUpdateRuntimeAppServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.CreateOrUpdateRuntimeAppService has not yet been implemented")
	})
	api.RuntimeAppCreateOrUpdateRuntimeAppWorkloadHandler = runtime_app.CreateOrUpdateRuntimeAppWorkloadHandlerFunc(func(params runtime_app.CreateOrUpdateRuntimeAppWorkloadParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.CreateOrUpdateRuntimeAppWorkload has not yet been implemented")
	})
	api.RuntimeAppsCreateRuntimeAppHandler = runtime_apps.CreateRuntimeAppHandlerFunc(func(params runtime_apps.CreateRuntimeAppParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_apps.CreateRuntimeApp has not yet been implemented")
	})
	api.RuntimeAppCreateRuntimeAppSecretHandler = runtime_app.CreateRuntimeAppSecretHandlerFunc(func(params runtime_app.CreateRuntimeAppSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.CreateRuntimeAppSecret has not yet been implemented")
	})
	api.BuildAppsDeleteBuildAppHandler = build_apps.DeleteBuildAppHandlerFunc(func(params build_apps.DeleteBuildAppParams) middleware.Responder {
		return middleware.NotImplemented("operation build_apps.DeleteBuildApp has not yet been implemented")
	})
	api.GroupsDeleteGroupHandler = groups.DeleteGroupHandlerFunc(func(params groups.DeleteGroupParams) middleware.Responder {
		return middleware.NotImplemented("operation groups.DeleteGroup has not yet been implemented")
	})
	api.GroupRuntimeDeleteGroupPVCHandler = group_runtime.DeleteGroupPVCHandlerFunc(func(params group_runtime.DeleteGroupPVCParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.DeleteGroupPVC has not yet been implemented")
	})
	api.GroupRuntimeDeleteGroupSecretHandler = group_runtime.DeleteGroupSecretHandlerFunc(func(params group_runtime.DeleteGroupSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.DeleteGroupSecret has not yet been implemented")
	})
	api.RuntimeAppsDeleteRuntimeAppHandler = runtime_apps.DeleteRuntimeAppHandlerFunc(func(params runtime_apps.DeleteRuntimeAppParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_apps.DeleteRuntimeApp has not yet been implemented")
	})
	api.RuntimeAppDeleteRuntimeAppSecretHandler = runtime_app.DeleteRuntimeAppSecretHandlerFunc(func(params runtime_app.DeleteRuntimeAppSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.DeleteRuntimeAppSecret has not yet been implemented")
	})
	api.RuntimeAppExecRuntimeAppPodHandler = runtime_app.ExecRuntimeAppPodHandlerFunc(func(params runtime_app.ExecRuntimeAppPodParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.ExecRuntimeAppPod has not yet been implemented")
	})
	api.BuildAppsGetBuildAppHandler = build_apps.GetBuildAppHandlerFunc(func(params build_apps.GetBuildAppParams) middleware.Responder {
		return middleware.NotImplemented("operation build_apps.GetBuildApp has not yet been implemented")
	})
	api.BuildAppsGetBuildAppsHandler = build_apps.GetBuildAppsHandlerFunc(func(params build_apps.GetBuildAppsParams) middleware.Responder {
		return middleware.NotImplemented("operation build_apps.GetBuildApps has not yet been implemented")
	})
	api.BuildAppGetBuildLogHandler = build_app.GetBuildLogHandlerFunc(func(params build_app.GetBuildLogParams) middleware.Responder {
		return middleware.NotImplemented("operation build_app.GetBuildLog has not yet been implemented")
	})
	api.BuildAppGetBuildPublishmentsHandler = build_app.GetBuildPublishmentsHandlerFunc(func(params build_app.GetBuildPublishmentsParams) middleware.Responder {
		return middleware.NotImplemented("operation build_app.GetBuildPublishments has not yet been implemented")
	})
	api.BuildAppGetBuildsHandler = build_app.GetBuildsHandlerFunc(func(params build_app.GetBuildsParams) middleware.Responder {
		return middleware.NotImplemented("operation build_app.GetBuilds has not yet been implemented")
	})
	api.ClusterGetClusterComponentsStatusHandler = cluster.GetClusterComponentsStatusHandlerFunc(func(params cluster.GetClusterComponentsStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation cluster.GetClusterComponentsStatus has not yet been implemented")
	})
	api.ClusterGetClusterResourceUsageHandler = cluster.GetClusterResourceUsageHandlerFunc(func(params cluster.GetClusterResourceUsageParams) middleware.Responder {
		return middleware.NotImplemented("operation cluster.GetClusterResourceUsage has not yet been implemented")
	})
	api.GroupsGetGroupHandler = groups.GetGroupHandlerFunc(func(params groups.GetGroupParams) middleware.Responder {
		return middleware.NotImplemented("operation groups.GetGroup has not yet been implemented")
	})
	api.GroupRuntimeGetGroupIngressHandler = group_runtime.GetGroupIngressHandlerFunc(func(params group_runtime.GetGroupIngressParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.GetGroupIngress has not yet been implemented")
	})
	api.GroupRuntimeGetGroupPVCHandler = group_runtime.GetGroupPVCHandlerFunc(func(params group_runtime.GetGroupPVCParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.GetGroupPVC has not yet been implemented")
	})
	api.GroupRuntimeGetGroupPVCsHandler = group_runtime.GetGroupPVCsHandlerFunc(func(params group_runtime.GetGroupPVCsParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.GetGroupPVCs has not yet been implemented")
	})
	api.GroupRuntimeGetGroupSecretHandler = group_runtime.GetGroupSecretHandlerFunc(func(params group_runtime.GetGroupSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.GetGroupSecret has not yet been implemented")
	})
	api.GroupRuntimeGetGroupSecretsHandler = group_runtime.GetGroupSecretsHandlerFunc(func(params group_runtime.GetGroupSecretsParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.GetGroupSecrets has not yet been implemented")
	})
	api.GroupsGetGroupsHandler = groups.GetGroupsHandlerFunc(func(params groups.GetGroupsParams) middleware.Responder {
		return middleware.NotImplemented("operation groups.GetGroups has not yet been implemented")
	})
	api.RuntimeAppsGetRuntimeAppHandler = runtime_apps.GetRuntimeAppHandlerFunc(func(params runtime_apps.GetRuntimeAppParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_apps.GetRuntimeApp has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppConfigMapHandler = runtime_app.GetRuntimeAppConfigMapHandlerFunc(func(params runtime_app.GetRuntimeAppConfigMapParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppConfigMap has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppConfigMapHistoryHandler = runtime_app.GetRuntimeAppConfigMapHistoryHandlerFunc(func(params runtime_app.GetRuntimeAppConfigMapHistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppConfigMapHistory has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppPodHandler = runtime_app.GetRuntimeAppPodHandlerFunc(func(params runtime_app.GetRuntimeAppPodParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppPod has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppPodsHandler = runtime_app.GetRuntimeAppPodsHandlerFunc(func(params runtime_app.GetRuntimeAppPodsParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppPods has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppSecretHandler = runtime_app.GetRuntimeAppSecretHandlerFunc(func(params runtime_app.GetRuntimeAppSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppSecret has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppSecretHistoryHandler = runtime_app.GetRuntimeAppSecretHistoryHandlerFunc(func(params runtime_app.GetRuntimeAppSecretHistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppSecretHistory has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppSecretsHandler = runtime_app.GetRuntimeAppSecretsHandlerFunc(func(params runtime_app.GetRuntimeAppSecretsParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppSecrets has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppServiceHandler = runtime_app.GetRuntimeAppServiceHandlerFunc(func(params runtime_app.GetRuntimeAppServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppService has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppServiceHistoryHandler = runtime_app.GetRuntimeAppServiceHistoryHandlerFunc(func(params runtime_app.GetRuntimeAppServiceHistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppServiceHistory has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppWorkloadHandler = runtime_app.GetRuntimeAppWorkloadHandlerFunc(func(params runtime_app.GetRuntimeAppWorkloadParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppWorkload has not yet been implemented")
	})
	api.RuntimeAppGetRuntimeAppWorkloadHistoryHandler = runtime_app.GetRuntimeAppWorkloadHistoryHandlerFunc(func(params runtime_app.GetRuntimeAppWorkloadHistoryParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.GetRuntimeAppWorkloadHistory has not yet been implemented")
	})
	api.RuntimeAppsGetRuntimeAppsHandler = runtime_apps.GetRuntimeAppsHandlerFunc(func(params runtime_apps.GetRuntimeAppsParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_apps.GetRuntimeApps has not yet been implemented")
	})
	api.RuntimeAppLogsRuntimeAppPodHandler = runtime_app.LogsRuntimeAppPodHandlerFunc(func(params runtime_app.LogsRuntimeAppPodParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.LogsRuntimeAppPod has not yet been implemented")
	})
	api.PingPingHandler = ping.PingHandlerFunc(handlers.Ping)
	api.BuildAppsUpdateBuildAppHandler = build_apps.UpdateBuildAppHandlerFunc(func(params build_apps.UpdateBuildAppParams) middleware.Responder {
		return middleware.NotImplemented("operation build_apps.UpdateBuildApp has not yet been implemented")
	})
	api.GroupsUpdateGroupHandler = groups.UpdateGroupHandlerFunc(func(params groups.UpdateGroupParams) middleware.Responder {
		return middleware.NotImplemented("operation groups.UpdateGroup has not yet been implemented")
	})
	api.GroupRuntimeUpdateGroupPVCHandler = group_runtime.UpdateGroupPVCHandlerFunc(func(params group_runtime.UpdateGroupPVCParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.UpdateGroupPVC has not yet been implemented")
	})
	api.GroupRuntimeUpdateGroupSecretHandler = group_runtime.UpdateGroupSecretHandlerFunc(func(params group_runtime.UpdateGroupSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation group_runtime.UpdateGroupSecret has not yet been implemented")
	})
	api.RuntimeAppsUpdateRuntimeAppHandler = runtime_apps.UpdateRuntimeAppHandlerFunc(func(params runtime_apps.UpdateRuntimeAppParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_apps.UpdateRuntimeApp has not yet been implemented")
	})
	api.RuntimeAppUpdateRuntimeAppSecretHandler = runtime_app.UpdateRuntimeAppSecretHandlerFunc(func(params runtime_app.UpdateRuntimeAppSecretParams) middleware.Responder {
		return middleware.NotImplemented("operation runtime_app.UpdateRuntimeAppSecret has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
