package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type GiphyGif struct {
	Type                     string
	Id                       string
	URL                      string
	Tags                     string
	BitlyGifURL              string `json:"bitly_gif_url"`
	ImageOriginalURL         string `json:"image_original_url"`
	ImageFixedHeightStillUrl string `json:"image_fixed_height_still_url"`
	ImageFixedHeightWidth    string `json:"image_fixed_height_width"`
	ImageFixedHeightHeight   string `json:"image_fixed_height_height"`
	ImageFixedHeightURL      string `json:"image_fixed_height_url"`
}

type GiphyAPIResponse struct {
	Data  []GiphyGif
	Query string
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	search := req.FormValue("q")
	url := fmt.Sprintf("http://giphy.com/api/gifs?tag=%s", url.QueryEscape(search))
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	giphyResponse := GiphyAPIResponse{}
	if err := json.Unmarshal(body, &giphyResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	giphyResponse.Query = search
	t := template.Must(template.ParseFiles("search.html"))
	if err := t.Execute(w, giphyResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
