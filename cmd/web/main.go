package main

import (
	"database/sql"
	"flag"
	"github.com/gonesoft/snippetbox/pkg/models/postgres"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")

	dns := flag.String("dns", "postgresql://postgres:password@localhost/snippetbox?sslmode=disable", "Postgres datasource name")

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//db.SetMaxOpenConns(25)
	//db.SetMaxIdleConns(2)

	svc := postgres.NewSnippetModel(db)

	app := &application{
		errorLog: errorLog,
		infoLog:  infolog,
		snippets: svc,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
