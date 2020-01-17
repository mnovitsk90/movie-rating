CX CLOUD GO ASSIGNMENT
===

Use OMDB API to check Rotten Tomatoes Ratings of user-supplied movies

+ OMDB API apiKey stored at ./go/.apiKeyStore
+ Dockerfile stored at ./go/Dockerfile
+ GoLang application file stored at ./go/movierating.go
+ Executable file stored at ./go/movierating.exe

RUN
---
```
$ go build ./go/movierating.go
$ ./movierating.exe <"title">
```

Example
---

```
.\go\movierating.exe "The Dark Knight"
```

---
Expected Output:

```
Starting execution
------------------
Searching for movie: The Dark Knight
------------------
Movie Rating Source: Rotten Tomatoes
Movie Rating: 94%
```