// main.go
// Snippetbox Application 12-24-24

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	snippets *SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "foo:bar@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{
		snippets: &SnippetModel{DB: db},
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/view/{id}", app.viewHandler)
	http.HandleFunc("/create", app.createHandler)
    	http.HandleFunc("/edit/{id}", app.editHandler)
    	http.HandleFunc("/delete/{id}", app.deleteHandler)

	log.Printf("starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(http.DefaultServeMux)))
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
