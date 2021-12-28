# urlshortner

This repository contains the API to generate short URL of the long input URL.
## Clone the project

```
$ git clone git@github.com:Pravin1989/urlshortner.git

```

## Two ways to run it 

```
1. If you have a Docker installed on your machine then run below commands
$ cd urlshortner
$ docker compose build
$ docker compose up
```
```
2. If you have installed only Go on your machine then run below commands
$ cd urlshortner
$ go build .\src\
$ run the .exe file
```

## API Details
* Endpoint : http://localhost:8090/urlshortner/api/create
* Http Method Type : **POST**
* Input Payload : `{"url":"https://test.com/generate/generate-tinyurl"}`
* Sample Response : `"http://urlshortner.com/dmNRbq"`
