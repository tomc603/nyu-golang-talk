package main

import (
	"html/template"
	"log"
	"net/http"
)

type student struct {
	Name     string
	ImageURL string
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))

	s := student{
		Name:     "Skids",
		ImageURL: "https://sphotos-a.xx.fbcdn.net/hphotos-ash3/222314_1959428265669_121419848_n.jpg",
	}
	if err := t.Execute(w, s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
