// models.go

package main

import (
	"database/sql"
	"log"
)

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created string
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) insert(title string, content string) (int, error) {
	stmt := "INSERT INTO snippets (title, content) VALUES(?, ?);"

	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		log.Println(err)
	}

	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) getOneSnippet(id int) (Snippet, error) {
	stmt := "SELECT * FROM snippets WHERE id = ?;"

	row := m.DB.QueryRow(stmt, id)
	// Initialize a new zeroed Snippet struct.
	var snippet Snippet

	err := row.Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Created)
	if err != nil {
		log.Println(err)
	}

	return snippet, nil
}

func (m *SnippetModel) getAllSnippets() ([]Snippet, error) {

	stmt := "SELECT * FROM snippets;"

	// Use the Query() method on the connection pool to execute our
	// SQL statement. This returns a sql.Rows resultset containing the result of
	// our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice of Snippet structs
	var snippets []Snippet

	for rows.Next() {
		var snippet Snippet

		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Created)
		if err != nil {
			return nil, err
		}

		// Append it to the slice of snippets.
		snippets = append(snippets, snippet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return snippets, nil
}

func (m *SnippetModel) update(id int, title string, content string) error {
	log.Println("You are in the update model.")
	stmt := "UPDATE snippets SET title=?, content=? WHERE id=?"
	_, err := m.DB.Exec(stmt, title, content, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *SnippetModel) delete(id int) error {
	log.Println("You are in the delete model.")
	stmt := "DELETE FROM snippets WHERE id=?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
