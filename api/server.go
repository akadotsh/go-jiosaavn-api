package api

import (
	"net/http"
	"time"

	"github.com/akadotsh/go-jiosaavn-client/api/middleware"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router = router.PathPrefix("/api").Subrouter()

	router.Use(middleware.Logging)
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

	srv := &http.Server{
		Handler: router,
		Addr:    s.listenAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}


