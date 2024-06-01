package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/akadotsh/go-jiosaavn-client/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	router.Use(loggingMiddleware)
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/songs/{id}", getSongByID).Methods("GET")
	router.HandleFunc("/songs/{id}/suggestions", getSongSuggestions).Methods("GET")
	router.HandleFunc("/songs/{id}/lyrics", getSongLyrics).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	router.HandleFunc("/search", searchAll).Methods("GET")
	router.HandleFunc("/search/songs", searchSongs).Methods("GET")
	router.HandleFunc("/search/albums", searchAlbums).Methods("GET")
	router.HandleFunc("/search/artists", searchArtists).Methods("GET")
	router.HandleFunc("/search/playlist", searchPlaylists).Methods("GET")
	router.HandleFunc("/artists/{id}", getArtistById).Methods("GET")
	router.HandleFunc("/artists/{id}/songs", getArtistSongs).Methods("GET")
	router.HandleFunc("/artists/{id}/albums", getArtistAlbums).Methods("GET")
	router.HandleFunc("/playlists", getPlaylistById).Methods("GET")

	return http.ListenAndServe(s.listenAddr, router)
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

func getArtistById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response := utils.FetchReq(utils.Artists.ID, "web6dot0", utils.Params{Key: "artistId", Value: id})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func getArtistAlbums(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var params []utils.Params

	for key, val := range r.URL.Query() {
		params = append(params, utils.Params{Key: key, Value: val[0]})
	}

	params = append(params, utils.Params{Key: "artistId", Value: id})

	fmt.Println("params", params)

	response := utils.FetchReq(utils.Artists.Albums, "web6dot0", params...)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func getArtistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var params []utils.Params

	for key, val := range r.URL.Query() {
		params = append(params, utils.Params{Key: key, Value: val[0]})
	}

	params = append(params, utils.Params{Key: "artistId", Value: id})

	response := utils.FetchReq(utils.Artists.Songs, "web6dot0", params...)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func getPlaylistById(w http.ResponseWriter, r *http.Request) {

	response := utils.FetchReq(utils.Playlist.ID, "web6dot0", utils.Params{Key: "listid", Value: r.URL.Query()["id"][0]})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
