package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

var ports = []string{"8000", "8080"}

func serve(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	addr := ctx.Value(http.LocalAddrContextKey).(net.Addr)
	fmt.Fprintf(w, "[%v] Hey there.\n", addr)
}

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	http.HandleFunc("/", serve)
	if len(os.Args) > 1 {
		ports = os.Args[1:]
	}
	for _, port := range ports {
		go func(port string) {
			log.Printf("Listening on %v\n", port)
			log.Fatal(http.ListenAndServe(":"+port, nil))
		}(port)
	}
	<-sig
	log.Println("Exiting.")
}
