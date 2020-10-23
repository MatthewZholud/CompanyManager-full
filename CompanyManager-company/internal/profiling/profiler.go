package profiling

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-company/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func ProfilingServer()  {
	r := mux.NewRouter()
	RegisterProfilingRoutes(r)
	err := http.ListenAndServe(":8006", r)
	if err != nil {
		logger.Log.Fatal("Can't connect to profiling server:", err)
	} else {
		logger.Log.Info("Connected to profiling server")
	}
}


// go tool pprof -http=":8000"  http://localhost:8006/debug/pprof/profile?seconds=5  to run profiling