package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
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
	logger := func(inner http.Handler, name string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			w.Header().Set("Access-Control-Allow-Origin", "*")
			inner.ServeHTTP(w, r)
			log.Info(fmt.Sprintf("served %s", name), "method", r.Method, "url", r.RequestURI, "duration", time.Since(start))
		})
	}

	for _, route := range buildRoutes(api) {
		var handler http.Handler = route.HandlerFunc
		handler = logger(handler, route.Name)

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
			"Ping",
			http.MethodGet,
			"/ping",
			api.Ping,
		},
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
