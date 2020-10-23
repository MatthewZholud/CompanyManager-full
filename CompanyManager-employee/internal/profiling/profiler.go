package profiling

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-employee/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func ProfilingServer()  {
	r := mux.NewRouter()
	RegisterProfilingRoutes(r)
	err := http.ListenAndServe(":8007", r)
	if err != nil {
		logger.Log.Fatal("Can't connect to profiling server:", err)
	} else {
		logger.Log.Info("Connected to profiling server")
	}
}
