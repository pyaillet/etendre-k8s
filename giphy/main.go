package main

import (
	"encoding/base64"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getGif(tag string, key string) Gif {
	img, err := ioutil.ReadFile("./cat.gif")
	check(err)
	b64 := base64.StdEncoding.EncodeToString(img)
	gif := Gif{Title: "cat", Img: b64, Tag: tag}
	return gif
}

func getTemplate() string {
	tmpl, err := ioutil.ReadFile("./template.html")
	check(err)
	return string(tmpl)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tag := os.Getenv("TAG")
	key := os.Getenv("GIPHY_APIKEY")
	html := getTemplate()
	gif := getGif(tag, key)
	t, err := template.New("html").Parse(html)
	check(err)
	err = t.Execute(w, gif)
	check(err)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
