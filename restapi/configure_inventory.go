// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"inventory-management/restapi/operations"
	"inventory-management/restapi/operations/authorization"
	"inventory-management/restapi/operations/stock"
	"inventory-management/restapi/operations/user"
	"inventory-management/restapi/pkg/controllers"
	"inventory-management/restapi/pkg/database"
)

//go:generate swagger generate server --target ../../inventory-management --name Inventory --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.InventoryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.InventoryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	
	database.Connect()
	
	api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams) middleware.Responder {
		controllers.CreateUser(params)
		// return middleware.ResponderFunc(func(rw http.ResponseWriter, p runtime.Producer) {
		// 	rw.Write([]byte("user created"))
		// })
		return user.NewPostUserCreated()
		//return middleware.NotImplemented("Data saved")
	})

	if api.StockDeleteInventoryItemIDHandler == nil {
		api.StockDeleteInventoryItemIDHandler = stock.DeleteInventoryItemIDHandlerFunc(func(params stock.DeleteInventoryItemIDParams) middleware.Responder {
			return middleware.NotImplemented("operation stock.DeleteInventoryItemID has not yet been implemented")
		})
	}
	if api.UserDeleteUserIDHandler == nil {
		api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserID has not yet been implemented")
		})
	}
	if api.StockGetInventoryHandler == nil {
		api.StockGetInventoryHandler = stock.GetInventoryHandlerFunc(func(params stock.GetInventoryParams) middleware.Responder {
			return middleware.NotImplemented("operation stock.GetInventory has not yet been implemented")
		})
	}
	if api.StockGetInventorySearchItemNameHandler == nil {
		api.StockGetInventorySearchItemNameHandler = stock.GetInventorySearchItemNameHandlerFunc(func(params stock.GetInventorySearchItemNameParams) middleware.Responder {
			return middleware.NotImplemented("operation stock.GetInventorySearchItemName has not yet been implemented")
		})
	}
	if api.UserGetUserIDHandler == nil {
		api.UserGetUserIDHandler = user.GetUserIDHandlerFunc(func(params user.GetUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserID has not yet been implemented")
		})
	}
	if api.StockPostInventoryHandler == nil {
		api.StockPostInventoryHandler = stock.PostInventoryHandlerFunc(func(params stock.PostInventoryParams) middleware.Responder {
			return middleware.NotImplemented("operation stock.PostInventory has not yet been implemented")
		})
	}
	if api.AuthorizationPostLoginHandler == nil {
		api.AuthorizationPostLoginHandler = authorization.PostLoginHandlerFunc(func(params authorization.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization.PostLogin has not yet been implemented")
		})
	}
	if api.StockPutInventoryItemIDHandler == nil {
		api.StockPutInventoryItemIDHandler = stock.PutInventoryItemIDHandlerFunc(func(params stock.PutInventoryItemIDParams) middleware.Responder {
			return middleware.NotImplemented("operation stock.PutInventoryItemID has not yet been implemented")
		})
	}
	if api.UserPutUserHandler == nil {
		api.UserPutUserHandler = user.PutUserHandlerFunc(func(params user.PutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	// create a new serve mux for handling additional urls
    mux := http.NewServeMux()

    // forward to go-swagger by default
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handler.ServeHTTP(w, r)
    })

	return handler
}
