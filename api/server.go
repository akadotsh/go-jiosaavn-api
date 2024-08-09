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

	router.HandleFunc("/",RootHandler)

	apiRouter:= router.PathPrefix("/api").Subrouter()

	apiRouter.Use(middleware.Logging)
	apiRouter.HandleFunc("/", RootHandler).Methods("GET")
	apiRouter.HandleFunc("/songs/{id}", getSongByID).Methods("GET")
	apiRouter.HandleFunc("/songs/{id}/suggestions", getSongSuggestions).Methods("GET")
	apiRouter.HandleFunc("/songs/{id}/lyrics", getSongLyrics).Methods("GET")
	apiRouter.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	apiRouter.HandleFunc("/search", searchAll).Methods("GET")
	apiRouter.HandleFunc("/search/songs", searchSongs).Methods("GET")
	apiRouter.HandleFunc("/search/albums", searchAlbums).Methods("GET")
	apiRouter.HandleFunc("/search/artists", searchArtists).Methods("GET")
	apiRouter.HandleFunc("/search/playlist", searchPlaylists).Methods("GET")
	apiRouter.HandleFunc("/artists/{id}", getArtistById).Methods("GET")
	apiRouter.HandleFunc("/artists/{id}/songs", getArtistSongs).Methods("GET")
	apiRouter.HandleFunc("/artists/{id}/albums", getArtistAlbums).Methods("GET")
	apiRouter.HandleFunc("/playlists", getPlaylistById).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    s.listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
