package main

import (
	"fmt"
	"net/http"
	"text/template"
	//"time"
)

type Login struct {
	MsgError   string
	IfLabelone string
	Labelone   string
}

func HandlerLoginHtml(w http.ResponseWriter, r *http.Request) {
	// template lendo HTML

	tmpl := template.Must(template.ParseFiles("login.html"))

	login := Login{
		MsgError:   "",
		IfLabelone: "",
		Labelone:   "logar!",
	}

	tmpl.Execute(w, login)
}

func HandlerAdminHtml(w http.ResponseWriter, r *http.Request) {

	// if session admin=1
}

func HandlerAuth(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(time.Second * 4)
	json := `{"status":"ok", "message":"tudo ocorreu bem na Auth..."}`
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func main() {

	mux := http.NewServeMux()

	// retorna HTML
	mux.HandleFunc("/login", HandlerLoginHtml)

	// auth
	mux.HandleFunc("/v1/api/auth", HandlerAuth)

	//mux.HandleFunc("/admin", HandlerAdminHtml)
	// mux.HandleFunc("/admin/user", Middleware(
	// 	jwtAut,
	// 	User,
	// ))
	//mux.HandleFunc("/admin/profile", HandlerAdminHtml)

	// fisico
	fs := http.FileServer(http.Dir("./web/static"))

	// vitual
	mux.Handle("/web/static/",
		http.StripPrefix("/web/static", fs))

	fmt.Println("Server Run: 8080")
	http.ListenAndServe(":8080", mux)
}
