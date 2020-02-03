# CX CLOUD - GoLang - Movie Rating Lookup

This service utilizes http://www.omdbapi.com/ to check the Rotten Tomatoes Rating associated with a given Movie Title.

+ Dockerfile stored at ./Dockerfile
+ GoLang application file stored at ./movieLookup.go
+ Executable file stored at ./getMovie

## Before Use:

1) Go to http://www.omdbapi.com/apikey.aspx in order to sign up for the service and acquire an API key.
2) Make sure you **set and export** a 'omdbApiKey' environment variable before running via CLI.
3) You **must** also pass the 'omdbApiKey' environment variable to the docker container when running the docker container.

## To Run:

## CLI:
```
$ export omdbApiKey='<your api key>'
$ go build -o getMovie ./movieLookup.go
$ ./getMovie [-t|--title] "<Movie Title>"
```

#### Example

```
$ ./getMovie -t "The Dark Knight"
94%

$ ./getMovie --title "The Dark Knight"
94%
```

## Docker:
```
$ export omdbApiKey='<your api key>'
$ docker build -t <image_name> .
$ docker run -e omdbApiKey=$omdbApiKey [-t|--title] "<Movie Title>"
```

#### Example

```
$ docker build -t movielookup:1.0 .

$ docker run -e omdbApiKey=$omdbApiKey movielookup:1.0 -t "The Dark Knight"
94%

$ docker run -e omdbApiKey=$omdbApiKey movielookup:1.0 --title "The Dark Knight"
94%
```