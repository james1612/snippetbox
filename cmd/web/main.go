package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)


	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal..
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	infoLog.Printf("Starting server on :9000")

	srv := &http.Server{
        Addr:     ":9000",
        ErrorLog: errorLog,
        Handler:  mux,
	}

	srv.ListenAndServe()

	err := http.ListenAndServe(":9000", mux)
	errorLog.Fatal(err)
}
