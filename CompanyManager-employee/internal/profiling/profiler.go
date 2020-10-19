package profiling

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ProfilingServer()  {
	r := mux.NewRouter()
	RegisterProfilingRoutes(r)
	log.Fatal(http.ListenAndServe(":8007", r))
}
