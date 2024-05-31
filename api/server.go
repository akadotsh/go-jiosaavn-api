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
	r.HandleFunc("/songs/{id}/lyrics", getSongLyrics).Methods("GET")
	r.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	r.HandleFunc("/search", searchAll).Methods("GET")
	r.HandleFunc("/search/songs", searchSongs).Methods("GET")
	r.HandleFunc("/search/albums", searchAlbums).Methods("GET")
	r.HandleFunc("/search/artists", searchArtists).Methods("GET")
	r.HandleFunc("/search/playlist", searchPlaylists).Methods("GET")

	return http.ListenAndServe(s.listenAddr, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("go saavan api"))
}

func getSongByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response := utils.FetchReq(utils.Songs.ID, "web6dot0", utils.Params{Key: "pids", Value: id})
	fmt.Println("response", response)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func getSongSuggestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	stationId := utils.FetchReq(utils.Songs.Station, "android", utils.Params{Key: "entity_id", Value: id}, utils.Params{Key: "entity_type", Value: "queue"})

	fmt.Println("stationId", stationId)

	response := utils.FetchReq(utils.Songs.Suggestions, "web6dot0", utils.Params{Key: "stationid", Value: "wL8VR46pohZAiLPopjqZT8tDsTIMnC-4swRpFLNuxcF7E-TF7reOA__"})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func getSongLyrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	response := utils.FetchReq(utils.Songs.Lyrics, "web6dot0", utils.Params{Key: "lyrics_id", Value: id})

	fmt.Println("response", response)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func getAlbumById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response := utils.FetchReq(utils.Album.ID, "web6dot0", utils.Params{Key: "albumid", Value: id})

	fmt.Println("response", response)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

// func getSongByLink(w http.ResponseWriter, r *http.Request){
//   vars:= mux.Vars(r);

//   link:= vars["link"]

//   response:=utils.FetchReq(utils.Songs.Link,"web6dot0",utils.Params{Key: "type",Value: "song"}, utils.Params{Key: "token",Value: link})

//   fmt.Println("response",response)

//   w.WriteHeader(http.StatusOK);

//   json.NewEncoder(w).Encode(response)
// }

func searchAll(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()["query"][0]

	response := utils.FetchReq(utils.Search.All, "web6dot0", utils.Params{Key: "query", Value: query})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func searchSongs(w http.ResponseWriter, r *http.Request) {
	params := utils.SearchParamBuilder(r.URL.Query())

	response := utils.FetchReq(utils.Search.Songs, "web6dot0", params...)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func searchAlbums(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response := utils.FetchReq(utils.Search.Songs, "web6dot0", params...)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func searchArtists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response := utils.FetchReq(utils.Search.Artists, "web6dot0", params...)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func searchPlaylists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response := utils.FetchReq(utils.Search.Playlists, "web6dot0", params...)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}
