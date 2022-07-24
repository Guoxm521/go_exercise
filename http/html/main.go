package main

import (
	"log"
	"net/http"
)

func main()  {
	dir :=	http.Dir("./root")
	staticHandler := http.FileServer(dir)
	http.Handle("/", http.StripPrefix("/", staticHandler))
	log.Printf("About to listen on 9090. Go to http://127.0.0.1:9090/")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}