package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func parseJsonToStruct(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return err
	}
	defer r.Body.Close()
	return nil
}

func IsNumericAndPositive(s string) bool {
	i, err := strconv.ParseFloat(s, 64)
	if err == nil && i >= 0 {
		return true
	} else {
		return false
	}
}

func Respond(w http.ResponseWriter, i interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(i)
	w.WriteHeader(http.StatusOK)
}
