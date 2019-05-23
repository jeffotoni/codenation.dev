package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

var (
	addr = ":443"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/api/ping",
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			io.WriteString(w, "DevopsBH, Golang for Devops TLS MUX!\n")
		})

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	// show
	log.Printf("Server Run %s TLS / https://localhost%s", addr, addr)

	// conf listen TLS
	err := srv.ListenAndServeTLS("server.crt", "server.key")
	log.Fatal(err)
}
