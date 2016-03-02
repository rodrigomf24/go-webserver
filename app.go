package main

import (
	"net/http"
	"html/template"
	"path"
	"fmt"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	http.ListenAndServe(":3333", nil)
}


func serveTemplate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URL PATH: %v", r.URL.Path);
	if r.URL.Path == "/" {
		r.URL.Path = "/main.html";
	} else if strings.Contains(r.URL.Path, ".html") == false {
		http.Redirect(w, r, "http://localhost:3333/", http.StatusFound);
		return
	}
	fmt.Println("URL PATH: %v", r.URL.Path);
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", r.URL.Path)

	tmpl, err := template.ParseFiles(lp, fp)

	if err != nil {
		http.Redirect(w, r, "http://localhost:3333/404.html", http.StatusFound);
		return
	}
	
	fmt.Println("ERROR: %v", err);

	tmpl.ExecuteTemplate(w, "layout", nil)
}