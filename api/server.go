package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"

	"github.com/akadotsh/go-jiosaavn-client/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Info(r.Method, r.RequestURI, time.Since(start))
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
	w.Write([]byte("Beep Boop!"))
}

type getSongByIDData struct {
	Modules map[string]any    `json:"modules"`
	Songs   []utils.SongsByID `json:"songs"`
}

func getSongByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.ID, "web6dot0", utils.Params{Key: "pids", Value: id})

	if err != nil {
		log.Error(err)
	}

	var data getSongByIDData
	json.Unmarshal(response, &data)

	if len(data.Songs) == 0 {
		msg := utils.ErrorHand{Status: "Failed", Message: "Song Not Found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data.Songs[0])
}

func getSongSuggestions(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	// id := vars["id"]

	// var stationIdRes,err = utils.FetchReq(utils.Songs.Station, "android", utils.Params{Key: "entity_id", Value: id}, utils.Params{Key: "entity_type", Value: "queue"})

	// fmt.Println("stationId", stationIdRes["stationid"])
	// stationId := stationIdRes["stationid"]

	// _id, ok := stationId.(string)

	// if !ok {
	// 	fmt.Println("Type assertion failed")
	// 	return
	// }

	// response,err := utils.FetchReq(utils.Songs.Suggestions, "android", utils.Params{Key: "stationid", Value: _id})

	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound);
	// 	json.NewEncoder(w).Encode("Artitst not Found")
	// }

	// w.WriteHeader(http.StatusOK)

	// json.NewEncoder(w).Encode(response)
}

func getSongLyrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.Lyrics, "web6dot0", utils.Params{Key: "lyrics_id", Value: id})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("lyrics not Found")
	}

	var data any
	json.Unmarshal(response, &data)
	fmt.Println("data", data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

func getAlbumById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Album.ID, "web6dot0", utils.Params{Key: "albumid", Value: id})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Albumn not Found")
	}

	var data any
	json.Unmarshal(response, &data)
	fmt.Println("data", data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func searchAll(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()["query"][0]

	response, err := utils.FetchReq(utils.Search.All, "web6dot0", utils.Params{Key: "query", Value: query})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No Results Found")
	}

	var data any
	json.Unmarshal(response, &data)
	fmt.Println("data", data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

func searchSongs(w http.ResponseWriter, r *http.Request) {
	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Songs, "web6dot0", params...)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No Results Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

func searchAlbums(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Songs, "web6dot0", params...)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Album not Found")
	}
	w.WriteHeader(http.StatusOK)

	var data any
	json.Unmarshal(response, &data)

	json.NewEncoder(w).Encode(data)

}

func searchArtists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Artists, "web6dot0", params...)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Artitst not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

func searchPlaylists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Playlists, "web6dot0", params...)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Playlist not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

func getArtistById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Artists.ID, "web6dot0", utils.Params{Key: "artistId", Value: id})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Artitst not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
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

	response, err := utils.FetchReq(utils.Artists.Albums, "web6dot0", params...)

	//TODO Later
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Artitst not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func getArtistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var params []utils.Params

	for key, val := range r.URL.Query() {
		params = append(params, utils.Params{Key: key, Value: val[0]})
	}

	params = append(params, utils.Params{Key: "artistId", Value: id})

	response, err := utils.FetchReq(utils.Artists.Songs, "web6dot0", params...)
	// TODO later
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Artitst not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func getPlaylistById(w http.ResponseWriter, r *http.Request) {

	response, err := utils.FetchReq(utils.Playlist.ID, "web6dot0", utils.Params{Key: "listid", Value: r.URL.Query()["id"][0]})

	//TODO later
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Artitst not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
