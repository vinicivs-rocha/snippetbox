package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLogger  *log.Logger
	infoLogger *log.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errLogger:  errorLogger,
		infoLogger: infoLogger,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	infoLogger.Printf("Listening on %s", *addr)
	err := server.ListenAndServe()
	errorLogger.Fatal(err)
}
