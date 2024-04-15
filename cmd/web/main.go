package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"snippetbox.vinicivs-rocha.com/internal/models"
)

type application struct {
	errLogger    *log.Logger
	infoLogger   *log.Logger
	snippetsRepo *models.SnippetRepository
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "snippetbox:123456@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)

	if err != nil {
		errorLogger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errLogger:    errorLogger,
		infoLogger:   infoLogger,
		snippetsRepo: &models.SnippetRepository{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	infoLogger.Printf("Listening on %s", *addr)
	err = server.ListenAndServe()
	errorLogger.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
