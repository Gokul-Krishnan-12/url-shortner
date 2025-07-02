package main

import (
	"log"
	"net/http"
)

func main()  {
	var err error
	urlStore, err = loadFromFile()
	if err != nil {
		log.Fatal("Failed to load store:", err)
	}

	// Serve static HTML files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))

	http.HandleFunc("/shorten",shortenUrl)
	http.HandleFunc("/",redirectUrl)
	http.HandleFunc("/stats/",statsHandler)

	log.Println("Starting server on : 9000")
	log.Fatal(http.ListenAndServe(":9000",nil))
}