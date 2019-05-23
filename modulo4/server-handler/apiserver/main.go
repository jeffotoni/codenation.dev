package main

import (
	"log"
	"net/http"
	"time"
)

// Adapter..
type Adapter func(http.Handler) http.Handler

func Use(next http.Handler, adapters ...Adapter) http.Handler {
	for _, adapterfunc := range adapters {
		next = adapterfunc(next)
	}

	return next
}

// nosso primeiro middleware
func Logger(name string) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %s %s\n",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}

func Ping1(w http.ResponseWriter, req *http.Request) {
	//server
	//log.Println("chamei minha api 1..")

	// client
	w.Write([]byte("\nDevops BH for Golang mux Handle()!"))
}

func Ping2(w http.ResponseWriter, req *http.Request) {
	//server
	//log.Println("chamei minha api 2..")

	// client
	w.Write([]byte("\nDevops BH for Golang mux Handle()!"))
}

func Ping3(w http.ResponseWriter, req *http.Request) {
	//server
	//log.Println("chamei minha api 3..")

	// client
	w.Write([]byte("\nDevops BH for Golang mux Handle()!"))
}
func main() {

	mux := http.NewServeMux()

	handlerPing1 := http.HandlerFunc(Ping1)
	handlerPing2 := http.HandlerFunc(Ping2)
	handlerPing3 := http.HandlerFunc(Ping3)

	// handlerFunc
	mux.Handle("/v1/api/ping1", Use(
		handlerPing1,
		Logger("Ping1"),
	)) // ok

	mux.Handle("/v1/api/ping2", Use(
		handlerPing2,
		Logger("Ping2"),
	)) // ok

	mux.Handle("/v1/api/ping3", Use(
		handlerPing3,
		Logger("Ping3"),
	)) // ok

	log.Printf("\nServer run 8082\n")
	// Listen
	log.Fatal(http.ListenAndServe(":8082", mux))
}
