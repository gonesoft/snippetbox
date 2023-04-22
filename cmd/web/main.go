package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/golangcollege/sessions"
	"github.com/gonesoft/snippetbox/pkg/models/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      database.SnippetModel
	templateCache map[string]*template.Template
	users         *database.UserModel
}

func main() {
	addr := flag.String("addr", ":8085", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	envErr := godotenv.Load()
	if envErr != nil {
		errorLog.Fatal("Error loading .env file: %v", envErr)
	}

	secret := os.Getenv("SESSION_KEY")
	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour
	session.SameSite = http.SameSiteStrictMode

	cfg := database.Config{
		Host:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSERNAME"),
		Password: os.Getenv("PGPASSWORD"),
		Name:     os.Getenv("PGDATABASE"),
		Port:     5432,
	}

	db, err := openDB(cfg)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	//db.SetMaxOpenConns(25)
	//db.SetMaxIdleConns(2)

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      database.SnippetModel{DB: db},
		templateCache: templateCache,
		users:         &database.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(cfg database.Config) (*sql.DB, error) {
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
