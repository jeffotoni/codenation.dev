package main

import (
	"html/template"
	"net/http"
)

// {{/* a comment */}}	Defines a comment
// {{.}}	Renders the root element
// {{.Title}}	Renders the “Title”-field in a nested element
// {{if .Done}} {{else}} {{end}}	Defines an if-Statement
// {{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
// {{block "content" .}} {{end}}	Defines a block with the name “content”

// <h1>{{.PageTitle}}<h1>
// <ul>
//     {{range .Todos}}
//         {{if .Done}}
//             <li class="done">{{.Title}}</li>
//         {{else}}
//             <li>{{.Title}}</li>
//         {{end}}
//     {{end}}
// </ul>

// data := TodoPageData{
// 	PageTitle: "My TODO list",
// 	Todos: []Todo{
// 		{Title: "Task 1", Done: false},
// 		{Title: "Task 2", Done: true},
// 		{Title: "Task 3", Done: true},
// 	},
// }

// func(w http.ResponseWriter, r *http.Request) {
// 	tmpl.Execute(w, "data goes here")
// }

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// tmpl, err := template.ParseFiles("layout.html")
	// or
	// tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list here...",
			Todos: []Todo{
				{Title: "Aula 1", Done: false},
				{Title: "Aula 2", Done: true},
				{Title: "Aula 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
