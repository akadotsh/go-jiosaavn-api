package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)


const (
 web6dot0= "web6dot0"
 android="android"
)

type ContextType string



func QueryBuilder(endpoint string, context ContextType) string{
	url:= url.URL{
		Scheme: "https",
		Host: "www.jiosaavn.com",
		Path: "api.php",
	  }
	
	if context == ""{
		context= web6dot0
	}  

	  queryParams:= url.Query();

	  queryParams.Add("__call",endpoint)
	  queryParams.Add("_format","json")
	  queryParams.Add("_marker","0");
	  queryParams.Add("api_version","4")
	  queryParams.Add("api_version","4")
	  queryParams.Add("ctx", string(context))
	  queryParams.Add("pids", "3IoDK8qI")
	
	
	  url.RawQuery = queryParams.Encode()  
	  
	  fmt.Println("url",url.String())

	  return url.String()

}

func main(){
	res:=fetchQuery()
	fmt.Println("Response:", res)
}

func fetchQuery() any {
	const endpoint="song.getDetails"
	url:= QueryBuilder(endpoint,"web6dot0")
		 
   
	 fmt.Print("url",url)
	 resp,err:=http.Get(url)
	 defer resp.Body.Close()
 
	 if err != nil{
		 panic(err)
	 }
 
	 body,err:=io.ReadAll(resp.Body)
   
	 if err != nil {
		 panic(err)
	 }
 
	 var data any
	 err = json.Unmarshal(body, &data)
	 if err != nil {
		 panic(err)
	 }

	 return data
 
}