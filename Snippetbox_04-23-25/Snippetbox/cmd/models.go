// models.go

package main

import (
	"database/sql"
	"log"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created string
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) insert(title string, content string) error {
	stmt := `INSERT INTO snippets(title, content) VALUES(?, ?)`
	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (m *SnippetModel) update(id int, title string, content string) error {
	stmt := "UPDATE snippets SET title=?, content=? WHERE id=?"
	_, err := m.DB.Exec(stmt, title, content, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *SnippetModel) delete(id int) error {
	stmt := "DELETE FROM snippets WHERE id=?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) getSnippet(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created FROM snippets WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	// Initialize a new zeroed Snippet struct.
	var snippet Snippet

	// Use row.Scan() to copy the values to the corresponding field in the Snippet struct.
	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created)
	if err != nil {
		log.Println(err)
	}

	return snippet, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) getAllSnippets() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created FROM snippets ORDER BY id DESC LIMIT 5`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Println(err)
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is always properly closed
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	var snippets []Snippet

	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Create a new zero-value Snippet struct.
		var s Snippet
		// Use rows.Scan() to copy the values from the row to the Snippet struct.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created)
		if err != nil {
			log.Println(err)
		}
		// Append it to the slice of snippets.
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
