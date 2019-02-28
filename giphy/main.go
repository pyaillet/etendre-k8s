package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Gif represents a gif base64 encoded img and its title
type Gif struct {
	Title string
	Img   string
	Tag   string
}

// GiphyResponse holds response data from Giphy
type GiphyResponse struct {
	Data GiphyData
}

// GiphyData holds response info from Giphy
// Only used parts are extracted
type GiphyData struct {
	ImageOriginalURL string `json:"image_original_url"`
	Title            string
}

const host = "api.giphy.com"
const urlTemplate = "https://%1s/v1/gifs/random?api_key=%2s&tag=%3s&rating=G"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getGif(tag string, key string) (*Gif, error) {
	u := fmt.Sprintf(urlTemplate, host, key, tag)
	var giphyResp GiphyResponse
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	rawResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	json.Unmarshal([]byte(rawResp), &giphyResp)
	gifResp, err := http.Get(giphyResp.Data.ImageOriginalURL)
	if err != nil {
		return nil, err
	}
	img, err := ioutil.ReadAll(gifResp.Body)
	if err != nil {
		return nil, err
	}
	gifResp.Body.Close()
	b64 := base64.StdEncoding.EncodeToString(img)
	return &Gif{Title: giphyResp.Data.Title, Img: b64, Tag: tag}, nil
}

func getTemplate() string {
	tmpl, err := ioutil.ReadFile("./template.html")
	check(err)
	return string(tmpl)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tag := os.Getenv("TAG")
	key := os.Getenv("GIPHY_APIKEY")
	log.Printf("Received query for tag %s", tag)
	html := getTemplate()
	gif, err := getGif(tag, key)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Error, service unavailable"))
		return
	}
	t, err := template.New("html").Parse(html)
	check(err)
	err = t.Execute(w, *gif)
	check(err)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Printf("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
