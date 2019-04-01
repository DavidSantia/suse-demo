package main

import (
	"fmt"
	"net/http"
        "os"

	"github.com/newrelic/go-agent"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<p>Santia's test website</p>")
	fmt.Fprintln(w, "<img src=\"/image/user.jpg\"></body></html>")
	fmt.Fprintln(w, "</body></html>")
}

func main() {
	var app newrelic.Application
	var err error

        config := newrelic.NewConfig("SUSE Go test", os.Getenv("NEW_RELIC_LICENSE_KEY"))
	if app, err = newrelic.NewApplication(config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", indexHandler))

	fs := http.FileServer(http.Dir("img/"))
	http.Handle("/image/", http.StripPrefix("/image/", fs))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
