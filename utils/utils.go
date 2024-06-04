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

type Query struct {
	endpoint string
	context  ContextType
	params   []Params
}

func QueryBuilder(endpoint string, context ContextType, params []Params) string {

	url := url.URL{
		Scheme: "https",
		Host:   "www.jiosaavn.com",
		Path:   "api.php",
	}

	if context == "" {
		context = web6dot0
	}

	queryParams := url.Query()

	queryParams.Add("__call", endpoint)
	queryParams.Add("_format", "json")
	queryParams.Add("_marker", "0")
	queryParams.Add("api_version", "4")
	queryParams.Add("api_version", "4")
	queryParams.Add("ctx", string(web6dot0))
	for _, param := range params {
		queryParams.Add(param.Key, param.Value)
	}

	url.RawQuery = queryParams.Encode()

	fmt.Println("url", url.String())
	return url.String()

}

func FetchReq(endpoint string, context ContextType, params ...Params) ([]byte, error) {

	url := QueryBuilder(endpoint, context, params)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Panic(err)

		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	fmt.Println("ERROR", err)
	if err != nil {
		log.Panic(err)
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
