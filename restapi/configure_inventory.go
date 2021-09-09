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

	// User API to Add, Delete, Get and Update
	api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams) middleware.Responder {
		newUser := controllers.CreateUser(params)

		return user.NewPostUserCreated().WithPayload(newUser)
	})

	api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams) middleware.Responder {
		err := controllers.DeleteUser(params)
		if err != nil {
			return user.NewDeleteUserIDNotFound()
		} else {
			return user.NewDeleteUserIDOK()
		}
	})

	api.UserGetUserIDHandler = user.GetUserIDHandlerFunc(func(params user.GetUserIDParams) middleware.Responder {
		userData := controllers.FindUserById(params)
		return user.NewGetUserIDOK().WithPayload(&userData)
	})

	api.UserPutUserHandler = user.PutUserHandlerFunc(func(params user.PutUserParams) middleware.Responder {
		updatedUser := controllers.UpdateUser(params)
		return user.NewPutUserCreated().WithPayload(updatedUser)
	})

	// Stock API to Create, Get, Patch, Delete, Search an Item
	api.StockPostInventoryHandler = stock.PostInventoryHandlerFunc(func(params stock.PostInventoryParams) middleware.Responder {
		postItem := controllers.CreateItem(params)
		return stock.NewPostInventoryCreated().WithPayload(postItem)
	})

	api.StockGetInventoryHandler = stock.GetInventoryHandlerFunc(func(params stock.GetInventoryParams) middleware.Responder {
		allItems := controllers.GetAllItems()
		return stock.NewGetInventoryOK().WithPayload(allItems)
	})

	api.StockPatchInventoryHandler = stock.PatchInventoryHandlerFunc(func(params stock.PatchInventoryParams) middleware.Responder {
		modifiedItem := controllers.UdateAnItem(params)
		return stock.NewPatchInventoryOK().WithPayload(modifiedItem)
	})

	api.StockDeleteInventoryItemIDHandler = stock.DeleteInventoryItemIDHandlerFunc(func(params stock.DeleteInventoryItemIDParams) middleware.Responder {
		err := controllers.DeleteAnItem(params)
		if err != nil {
			return stock.NewDeleteInventoryItemIDNotFound()
		} else {
			return stock.NewDeleteInventoryItemIDOK()
		}
	})

	api.StockGetInventorySearchItemNameHandler = stock.GetInventorySearchItemNameHandlerFunc(func(params stock.GetInventorySearchItemNameParams) middleware.Responder {
		return middleware.NotImplemented("operation stock.GetInventorySearchItemName has not yet been implemented")
	})

	//Auth

	api.AuthorizationPostLoginHandler = authorization.PostLoginHandlerFunc(func(params authorization.PostLoginParams) middleware.Responder {
		//Login not implemented
		controllers.LogInUser(params)
		return middleware.NotImplemented("operation authorization.PostLogin has not yet been implemented")
	})

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
