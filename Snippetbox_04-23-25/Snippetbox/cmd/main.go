// main.go

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependencies.
type application struct {
	snippets *SnippetModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "mike:5454160s@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// We pass openDB() the DSN from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// Call to db.Close() to close the connection pool before the main() function exits.
	defer db.Close()

	app := &application{
		snippets: &SnippetModel{DB: db},
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.indexHandler)
	mux.HandleFunc("/view/{id}", app.viewHandler)
	mux.HandleFunc("/create", app.createHandler)
	mux.HandleFunc("/edit/{id}", app.editHandler)
	mux.HandleFunc("/delete/{id}", app.deleteHandler)
	mux.HandleFunc("/about", app.aboutHandler)

	log.Printf("Starting HTTP Server on %s", *addr)
	err = http.ListenAndServe(*addr, logRequest(mux))
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
	os.Exit(1)
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
