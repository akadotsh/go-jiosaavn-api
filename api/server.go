package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/akadotsh/go-jiosaavn-client/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/songs/{id}", getSongByID).Methods("GET")
	r.HandleFunc("/songs/{id}/suggestions", getSongSuggestions).Methods("GET")

	return http.ListenAndServe(s.listenAddr, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("go saavan api"))
}

func getSongByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	resp := utils.FetchReq(utils.Songs.ID, "web6dot0", utils.Params{Key: "pids", Value: id})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func getSongSuggestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	fmt.Println("ID", id)
	resp := utils.FetchReq(utils.Songs.Suggestions, "web6dot0")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
