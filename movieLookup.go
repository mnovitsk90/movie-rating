package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "log"
  "net/url"
  "net/http"
  "encoding/json"
  "flag"
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

var inputMovie string

func init() {
  const (
    defaultMovie = ""
    usage = "Title of movie to lookup"
  )
  flag.StringVar(&inputMovie, "title", defaultMovie, usage)
  flag.StringVar(&inputMovie, "t", defaultMovie, usage)

  flag.Parse()
}

// func getMovieTitle() string {

//   inputMovie := "Boondock Saints"

//   if len(os.Args) > 1 {
//     inputMovie = os.Args[1]
//   } else {
//     inputMovie = "Boondock Saints"
//   }

//   return inputMovie
// }

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

  omdbApiKey, found := os.LookupEnv("omdbApiKey")

	if !found {
		log.Fatalf("Missing omdbApiKey env variable")
	}
  
	return omdbApiKey
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
      // fmt.Printf("Movie Rating Source: %s\n", movieObj.Ratings[i].Source)
      fmt.Printf("%s\n", movieObj.Ratings[i].Value)
			// fmt.Printf("Movie Rating: %s\n", movieObj.Ratings[i].Value)
		}
	}
}

func main() {

  // fmt.Println()
  // fmt.Println("Starting execution")
  // fmt.Println("------------------")

  fmt.Printf("%s\n", inputMovie)
  apiKey := getApiKey()
  // movieTitle := getMovieTitle()
  // queryUrl := buildUrl(movieTitle, apiKey)
  queryUrl := buildUrl(inputMovie, apiKey)
  
  // fmt.Printf("Searching for movie: %s\n", movieTitle)
  // fmt.Println("------------------")

  // fmt.Printf("Using this search URL: %q\n", queryUrl)

  movieInfo := getMovieInfo(queryUrl)

  // fmt.Println(string(movieInfo))
  
  getMovieRating(movieInfo) 

  // fmt.Println(body.Headers)
}
