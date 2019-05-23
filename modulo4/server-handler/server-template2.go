package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func IndexHtml(w http.ResponseWriter, req *http.Request) {

	// tmpl, err := template.ParseFiles("layout.html")
	// or
	// tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// <h1>{{.PageTitle}}<h1>
	// <ul>
	//     {{range .Todos}}
	//         {{if .Done}}
	//             <li class="done">{{.Title}}</li>
	//         {{else}}
	//             <li>{{.Title}}</li>
	//         {{end}}
	//     {{end}}
	// </ul

	data := TodoPageData{
		PageTitle: "My TODO list here...",
		Todos: []Todo{
			{Title: "Aula 1", Done: false},
			{Title: "Aula 2", Done: true},
			{Title: "Aula 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", IndexHtml)
	http.ListenAndServe(":8080", nil)
}
