package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type ContextType string

const (
	web6dot0 = "web6dot0"
	android  = "android"
)

type Params struct {
	Key   string
	Value string
}

func QueryBuilder(endpoint string, params []Params) string {

	url := url.URL{
		Scheme: "https",
		Host:   "www.jiosaavn.com",
		Path:   "api.php",
	}

	ctxValue := string(web6dot0)

	queryParams := url.Query()

	queryParams.Add("__call", endpoint)
	queryParams.Add("_format", "json")
	queryParams.Add("_marker", "0")
	queryParams.Add("api_version", "4")

	for _, param := range params {
		if param.Key == "ctx" {
			ctxValue = param.Value
		} else {
			queryParams.Add(param.Key, param.Value)

		}
	}
	queryParams.Add("ctx", ctxValue)

	url.RawQuery = queryParams.Encode()

	fmt.Println("url", url.String())
	return url.String()

}

func FetchReq(endpoint string, params ...Params) ([]byte, error) {

	url := QueryBuilder(endpoint, params)

	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println("ERROR", err)
	if err != nil {
		log.Panic(err)
		return nil, err

	}

	return body, nil
}

func FetchStationId(id []byte) ([]byte, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "www.jiosaavn.com",
		Path:   "api.php",
	}
	queryParams := url.Query()

	queryParams.Add("__call", Songs.Station)
	queryParams.Add("_format", "json")
	queryParams.Add("entity_id", string(id))
	queryParams.Add("entity_type", "queue")
	queryParams.Add("_marker", "0")
	queryParams.Add("api_version", "4")
	queryParams.Add("ctx", android)

	url.RawQuery = queryParams.Encode()

	resp, err := http.Get(url.String())

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func SearchParamBuilder(queries url.Values) []Params {
	var params []Params

	for key, Value := range queries {
		var val = Value[0]
		var k string
		if key == "query" {
			k = "q"
		} else if key == "page" {
			k = "p"
		} else {
			k = "n"
		}

		param := Params{Key: k, Value: val}

		params = append(params, param)
	}

	return params
}

type Response[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type GetSongByIDData struct {
	Modules map[string]any `json:"modules"`
	Songs   []SongsByID    `json:"songs"`
}

type SongsByIDMoreInfo struct {
	Music               string `json:"music"`
	Album_Id            string `json:"album_id"`
	Album               string `json:"album"`
	Label               string `json:"label"`
	Origin              string `json:"origin"`
	Is_Dolby_Contet     bool   `json:"is_dolby_content"`
	Encrypted_Media_Url string `json:"encrypted_media_url"`
	Album_Url           string `json:"album_url"`
	Duration            string `json:"duration"`
	Rights              struct {
		Code                 string `json:"code"`
		Cacheable            string `json:"cacheable"`
		Delete_Cached_Object string `json:"delete_cached_object"`
		Reason               string `json:"reason"`
	} `json:"rights"`
	Cache_Data     string `json:"cache_data"`
	Has_Lyrics     string `json:"has_lyrics"`
	Lyrics_Snippet string `json:"lyrics_support"`
	Starred        string `json:"starred"`
	Copyright_Text string `json:"copyright_text"`
	//	TODO   ArtistMap struct {
	//		Primary_Artists
	//	   }
	Release_Date         string `json:"release_date"`
	Label_Url            string `json:"label_url"`
	Vcode                string `json:"vcode"`
	Vlink                string `json:"vlink"`
	Triller_Available    bool   `json:"triller_available"`
	Request_Jiotune_Flag bool   `json:"request_jiotune_flag"`
	Webp                 string `json:"webp"`
	Lyrics_Id            string `json:"lyrics_id"`
}

type SongsByID struct {
	ID               string            `json:"id"`
	Title            string            `json:"title"`
	Subtitle         string            `json:"subtitle"`
	Header_Desc      string            `json:"header_desc"`
	Type             string            `json:"type"`
	Perma_Url        string            `json:"perma_url"`
	Image            string            `json:"image"`
	Play_Count       string            `json:"play_count"`
	Explicit_Content string            `json:"explicit_content"`
	List_Count       string            `json:"list_count"`
	List_Type        string            `json:"list_type"`
	List             string            `json:"list"`
	More_Info        SongsByIDMoreInfo `json:"more_info"`
}

type ErrorHand struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetAlbumByIdResponse struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	Subtitle         string      `json:"subtitle"`
	Header_Desc      string      `json:"header_desc"`
	Type             string      `json:"type"`
	Perma_Url        string      `json:"perma_url"`
	Image            string      `json:"image"`
	Language         string      `json:"language"`
	Year             string      `json:"year"`
	Play_Count       string      `json:"play_count"`
	Explicit_Content string      `json:"explicit_content"`
	List_Count       string      `json:"list_count"`
	List_Type        string      `json:"list_type"`
	List             []SongsByID `json:"list"`
	More_Info        struct {
		//TODO artistMap
		Song_Count       string `json:"song_count"`
		Copyright_Text   string `json:"copyright_text"`
		Is_Dolby_Content bool   `json:"is_dolby_content"`
		Label_Url        string `json:"label_url"`
	} `json:"more_info"`
}

type DownloadLink struct {
	Quality string `json:"quality"`
	Url     string `json:"url"`
}

type AlbumResponse struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Year            int            `json:"year"`
	Type            string         `json:"type"`
	PlayCount       int            `json:"playCount"`
	Language        string         `json:"language"`
	ExplicitContent bool           `json:"explicitContent"`
	Artists         any            `json:"artists"`
	SongCount       int            `json:"songCount"`
	Url             string         `json:"url"`
	Image           []DownloadLink `json:"image"`
	Songs           any            `json:"songs"`
}

type AlbumApiResponse struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Subtitle         string `json:"subtitle"`
	Header_desc      string `json:"header_desc"`
	Type             string `json:"type"`
	Perma_Url        string `json:"perma_url"`
	Image            string `json:"image"`
	Language         string `json:"language"`
	Year             string `json:"year"`
	Play_Count       string `json:"play_count"`
	Explicit_Content string `json:"explicit_content"`
	List_Count       string `json:"list_count"`
	List_Type        string `json:"list_type"`
	List             any    `json:"list"`
	More_Info        struct {
		artistMap        string
		song_count       string
		is_dolby_content bool
		label_url        string
	} `json:"more_info"`
}

// type SearchAllResponse struct {
// 	Albums struct{
// 		Data []struct{
// 			ID string `json:"id"`
// 			Title string `json:"title"`
// 			Subtitle string `json:"subtitle"`
// 			Type string `json:"type"`
// 			Image string `json:"image"`
// 			Perma_Url string `json:"perma_url"`
// 			More_Info struct{
// 				Music string `json:"music"`
// 				Ctr int `json:"ctr"`
// 				Year string `json:"year"`
// 				Is_Movie string `json:"is_movie"`
// 				Language string `json:"language"`
// 				Song_Pids string `json:"song_pids"`

// 			} `json:"more_info"`
// 		} `json:"data"`
// 	  Position int `json:"position"`
// 	} `json:"albums"`
//   Songs struct {
// 	Data []struct {
// 		ID string `json:"id"`
// 		Title string `json:"title"`

// 	}
//   }
// }

type LyrcisAPIResponse struct {
	Lyrics           string
	Lyrics_copyright string
	Snippet          string
}

type StationIdResponse struct {
	Stationid string
}
