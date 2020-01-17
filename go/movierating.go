package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "log"
  "net/url"
  "net/http"
  "encoding/json"
)

//Define Movie Data structure for HTTP GET return
type Movie struct {
	Ratings []Ratings `json:"Ratings"`
}

//Define Movie Ratings Data structure
type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

// json data type
var movieObj Movie

func getMovieTitle() string {

  inputMovie := "Boondock Saints"

  if len(os.Args) > 1 {
    inputMovie = os.Args[1]
  } else {
    fmt.Printf("No movie supplied, using default movie: %s\n", inputMovie)
  }

  return inputMovie
}

func buildUrl(movieTitle string, apiKey string) string {

  baseUrl, err := url.Parse("http://www.omdbapi.com/")

  if err != nil {
	  log.Fatalln(err)
  }

  params := url.Values{}
  params.Add("t", movieTitle)
  params.Add("apikey", apiKey)

  baseUrl.RawQuery = params.Encode()

  return baseUrl.String()
}

func getMovieInfo(fullUrl string) string {

  resp, err := http.Get(fullUrl)

  if err != nil {
	  log.Fatalln(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  return string(body)
}

func getApiKey() string {

  file, err := os.Open("./go/.apiKeyStore")
	if err != nil {
		fmt.Println(err)
  }
  
  defer file.Close()
  
  key, err := ioutil.ReadAll(file)
  if err != nil {
		fmt.Println(err)
  }
  
	return string(key)
}

func IsValidSource(source string) bool {
	switch source {
	case
		"Rotten Tomatoes":
		return true
	}
	return false
}

func getMovieRating(body string) {
  
  json.Unmarshal([]byte(body), &movieObj)

  for i := 0; i < len(movieObj.Ratings); i++ {

		if IsValidSource(movieObj.Ratings[i].Source) {
			fmt.Printf("Movie Rating Source: %s\n", movieObj.Ratings[i].Source)
			fmt.Printf("Movie Rating: %s\n", movieObj.Ratings[i].Value)
		}
	}
}

func main() {

  fmt.Println()
  fmt.Println("Starting execution")
  fmt.Println("------------------")

  apiKey := getApiKey()
  movieTitle := getMovieTitle()
  queryUrl := buildUrl(movieTitle, apiKey)
  
  fmt.Printf("Searching for movie: %s\n", movieTitle)
  fmt.Println("------------------")

  //fmt.Printf("Using this search URL: %q\n", queryUrl)

  movieInfo := getMovieInfo(queryUrl)

  //fmt.Println(string(movieInfo))
  
  getMovieRating(movieInfo) 

  //fmt.Println(body.Headers)
}
