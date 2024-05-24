package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
	}

	infoLog.Printf("Starting server on :9000")

	srv := &http.Server{
        Addr:     ":9000",
        ErrorLog: errorLog,
        Handler:  app.routes(),
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
