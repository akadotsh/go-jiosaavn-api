package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/akadotsh/go-jiosaavn-client/utils"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	resp := utils.Response{
		Status:  "success",
		Message: "Beep Boop!",
	}
	json.NewEncoder(w).Encode(resp)
}

type getSongByIDData struct {
	Modules map[string]any    `json:"modules"`
	Songs   []utils.SongsByID `json:"songs"`
}

func getSongByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.ID, utils.Params{Key: "pids", Value: id})

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

func getSongLyrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Songs.Lyrics, utils.Params{Key: "lyrics_id", Value: id})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("lyrics not Found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)

}

type StationIdResponse struct {
	Stationid string
}

func getSongSuggestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var encodedUriId, errP = json.Marshal([]string{url.QueryEscape(id)})

	if errP != nil {
		log.Error(errP)
	}

	var stationIdRes, err = utils.FetchStationId(encodedUriId)

	if err != nil {
		fmt.Println(err)
	}

	var stationId StationIdResponse
	json.Unmarshal(stationIdRes, &stationId)
	fmt.Println("stationId", stationId)

	response, err := utils.FetchReq(utils.Songs.Suggestions, utils.Params{Key: "ctx", Value: "android"}, utils.Params{Key: "stationid", Value: stationId.Stationid}, utils.Params{Key: "k", Value: "10"})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No song suggestion found")
	}

	var data any
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	response, err := utils.FetchReq(utils.Album.ID, utils.Params{Key: "albumid", Value: id})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Albumn not Found")
	}

	var data utils.GetAlbumByIdResponse
	json.Unmarshal(response, &data)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data.List)
}

func searchAll(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()["query"][0]

	response, err := utils.FetchReq(utils.Search.All, utils.Params{Key: "query", Value: query})

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

	response, err := utils.FetchReq(utils.Search.Songs, params...)
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

	response, err := utils.FetchReq(utils.Search.Songs, params...)
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

	response, err := utils.FetchReq(utils.Search.Artists, params...)
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

	response, err := utils.FetchReq(utils.Search.Playlists, params...)
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

	response, err := utils.FetchReq(utils.Artists.ID, utils.Params{Key: "artistId", Value: id})

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

	response, err := utils.FetchReq(utils.Artists.Albums, params...)

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

	response, err := utils.FetchReq(utils.Artists.Songs, params...)
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

	response, err := utils.FetchReq(utils.Playlist.ID, utils.Params{Key: "listid", Value: r.URL.Query()["id"][0]})

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

// https://www.jiosaavn.com/api.php?__call=webradio.createEntityStation&_format=json&_marker=0&api_version=4&ctx=android&entity_id=%5B%22yDeAS8Eh%22%5D&entity_type=queue
