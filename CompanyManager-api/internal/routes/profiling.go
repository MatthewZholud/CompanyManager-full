package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/pprof"
)

func RegisterProfilingRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(
		"/debug/pprof/", pprof.Index,
	)
	r.HandleFunc(
		"/debug/pprof/cmdline", pprof.Cmdline,
	)
	r.HandleFunc(
		"/debug/pprof/profile", pprof.Profile,
	)
	r.HandleFunc(
		"/debug/pprof/symbol", pprof.Symbol,
	)
	r.HandleFunc(
		"/debug/pprof/trace", pprof.Trace,
	)
	r.Handle(
		"/debug/pprof/goroutine", pprof.Handler("goroutine"),
	)
	r.Handle(
		"/debug/pprof/heap", pprof.Handler("heap"),
	)
	r.Handle(
		"/debug/pprof/threadcreate", pprof.Handler("threadcreate"),
	)
	r.Handle(
		"/debug/pprof/block", pprof.Handler("block"),
	)
	r.Handle(
		"/debug/vars", http.DefaultServeMux,
	)
	return r
}


// go tool pprof -http=":8000"  http://localhost:8006/debug/pprof/profile?seconds=5  to run profiling