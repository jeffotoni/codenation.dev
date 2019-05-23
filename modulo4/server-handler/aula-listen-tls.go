package main

// #generating .key and .csr
// $ openssl req -nodes -newkey rsa:2048 -keyout server.key -out server.csr -subj "/C=BR/ST=Minas/L=Belo Horizonte/O=s3wf Ltd./OU=IT/CN=localhost"

// # generating server .crt or .pem
// $ openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 365

import (
	"fmt"
	//"log"
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	//w.Write
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Ping na porta 443\n")
}

func main() {
	fmt.Println("Server Run:443")
	http.HandleFunc("/api/v1/ping", Ping) // != http.HandlerFunc()
	http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
