package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(api *Api) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range buildRoutes(api) {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func buildRoutes(api *Api) []Route {
	return []Route{
		{
			"SubmitUserOperation",
			http.MethodPost,
			"/userOperation",
			api.SubmitUserOperation,
		},
		{
			"GetSolverOperations",
			http.MethodGet,
			"/solverOperations",
			api.GetSolverOperations,
		},
		{
			"SubmitBundleOperations",
			http.MethodPost,
			"/bundleOperations",
			api.SubmitBundleOperations,
		},
		{
			"GetBundleHash",
			http.MethodGet,
			"/bundleHash",
			api.GetBundleHash,
		},
		{
			"SubmitSolverOperation",
			http.MethodPost,
			"/solverOperation",
			api.SubmitSolverOperation,
		},
		{
			"WebsocketSolver",
			http.MethodGet,
			"/ws/solver",
			api.WebsocketSolver,
		},
		{
			"WebsocketBundler",
			http.MethodGet,
			"/ws/bundler",
			api.WebsocketBundler,
		},
	}
}
