package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on Http Port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+ *port, nil))
}
