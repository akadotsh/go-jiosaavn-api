package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/akadotsh/go-jiosaavn-client/utils"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)


func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	
	handleSuccess(w, http.StatusOK, "Beep Boop!")
}

func getSongByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.ID, utils.Params{Key: "pids", Value: id})

	if err != nil {
		log.Error(err)
		handleErr(w, http.StatusInternalServerError, "Song Not Found")
		return
	}

	var data utils.GetSongByIDData
	json.Unmarshal(response, &data)

	if len(data.Songs) == 0 {
		handleErr(w, http.StatusNotFound, "Song Not Found")
		return
	}

	handleSuccess(w, http.StatusOK, data.Songs[0])

}

func getSongLyrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.Lyrics, utils.Params{Key: "lyrics_id", Value: id})

	if err != nil {

		handleErr(w, http.StatusNotFound, "lyrics not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func getSongSuggestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var encodedUriId, errP = json.Marshal([]string{url.QueryEscape(id)})

	if errP != nil {
		log.Error(errP)
		handleErr(w, http.StatusInternalServerError, errP.Error())
		return
	}

	var stationIdRes, err = utils.FetchStationId(encodedUriId)

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var stationId utils.StationIdResponse
	json.Unmarshal(stationIdRes, &stationId)

	response, err := utils.FetchReq(utils.Songs.Suggestions, utils.Params{Key: "ctx", Value: "android"}, utils.Params{Key: "stationid", Value: stationId.Stationid}, utils.Params{Key: "k", Value: "10"})

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Album.ID, utils.Params{Key: "albumid", Value: id})

	if err != nil || response == nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data utils.GetAlbumByIdResponse
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func searchAll(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()["query"][0]

	response, err := utils.FetchReq(utils.Search.All, utils.Params{Key: "query", Value: query})

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return

	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func searchSongs(w http.ResponseWriter, r *http.Request) {
	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Songs, params...)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func searchAlbums(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Songs, params...)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func searchArtists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Artists, params...)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func searchPlaylists(w http.ResponseWriter, r *http.Request) {

	params := utils.SearchParamBuilder(r.URL.Query())

	response, err := utils.FetchReq(utils.Search.Playlists, params...)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func getArtistById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Artists.ID, utils.Params{Key: "artistId", Value: id})

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func getArtistAlbums(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var params []utils.Params

	for key, val := range r.URL.Query() {
		params = append(params, utils.Params{Key: key, Value: val[0]})
	}

	params = append(params, utils.Params{Key: "artistId", Value: id})

	response, err := utils.FetchReq(utils.Artists.Albums, params...)

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func getArtistSongs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var params []utils.Params

	for key, val := range r.URL.Query() {
		params = append(params, utils.Params{Key: key, Value: val[0]})
	}

	params = append(params, utils.Params{Key: "artistId", Value: id})

	response, err := utils.FetchReq(utils.Artists.Songs, params...)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func getPlaylistById(w http.ResponseWriter, r *http.Request) {

	response, err := utils.FetchReq(utils.Playlist.ID, utils.Params{Key: "listid", Value: r.URL.Query()["id"][0]})

	if err != nil {
		handleErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	var data any
	json.Unmarshal(response, &data)

	handleSuccess(w, http.StatusOK, data)

}

func handleSuccess[T any](w http.ResponseWriter, status int, data T) {
	w.WriteHeader(status)
	resp := utils.Response[T]{
		Status: "success",
		Data:   data,
	}

	json.NewEncoder(w).Encode(resp)

}

func handleErr(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	resp := utils.Response[string]{
		Status: "Failed",
		Data:   errMsg,
	}
	json.NewEncoder(w).Encode(resp)
}
