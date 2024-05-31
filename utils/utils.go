package utils

import (
	"encoding/json"
	"fmt"
	"io"
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
	queryParams.Add("ctx", string(context))
	for _, param := range params {
		queryParams.Add(param.Key, param.Value)
	}

	url.RawQuery = queryParams.Encode()

	fmt.Println("url", url.String())

	return url.String()

}

func FetchReq(endpoint string, context ContextType, params ...Params) any {

	url := QueryBuilder(endpoint, context, params)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var data map[string]any
	json.Unmarshal(body, &data)

	return data
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
