# Giphy random tag server

## Build

`make`

## Run

`docker container run -p 8080:8080 -e TAG=cat -e GIPHY_APIKEY=<KEY> pyaillet/giphyserver:0.1`
