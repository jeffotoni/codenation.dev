package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func Ping1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if strings.ToUpper(r.Method) == "POST" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))
}

func Ping2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json := `{"status":"ok", "msg":"sucesso"}`
	json += "\n"
	w.Write([]byte(json))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/ping1", Ping1)
	mux.Handle("/api/v1/ping2", http.HandlerFunc(Ping2))

	confserver := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 3,
	}

	log.Println("Server run:8080")
	// http.ListenAndServe(":8080", mux)
	log.Fatal(confserver.ListenAndServe())

	var errs = make(chan error, 2)

	// async
	go func() {
		csi := make(chan os.Signal)
		signal.Notify(csi, os.Interrupt)
		errs <- fmt.Errorf("nosso notify.. %s", <-csi)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutdown Server!!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second*5)
	defer cancel()

	if err := confserver.Shutdown(ctx); err != nil {
		err = confserver.Close()
		if err != nil {
			log.Println("Shudown Server now...", err)
		}
	}

	log.Println("Server Exit...")
	<-errs

}
