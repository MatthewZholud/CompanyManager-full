package profiling

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ProfilingServer()  {
	r := mux.NewRouter()
	RegisterProfilingRoutes(r)
	log.Fatal(http.ListenAndServe(":8006", r))
}


// go tool pprof -http=":8000"  http://localhost:8006/debug/pprof/profile?seconds=5  to run profiling