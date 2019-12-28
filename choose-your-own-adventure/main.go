package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Option struct {
	Text        string `json:"text"`
	ChapterName string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Adventure map[string]Chapter

type AdventureHandler struct {
	adventure Adventure
	tpl       string
}

func (h AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles(h.tpl))
	chapterName := r.URL.Query().Get("chaptername")

	tpl.Execute(w, h.adventure[chapterName])
}

func main() {
	port := flag.Int("port", 8080, "port to listen the application on")
	jsonFilename := flag.String("json", "adventure.json", "path to the adventure json")
	HTMLtemplate := flag.String("template", "template.html", "path to the html template")
	flag.Parse()

	var handler AdventureHandler

	file, _ := ioutil.ReadFile(*jsonFilename)
	json.Unmarshal(file, &handler.adventure)
	handler.tpl = *HTMLtemplate

	http.Handle("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
