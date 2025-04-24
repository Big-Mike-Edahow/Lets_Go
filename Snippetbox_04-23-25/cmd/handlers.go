// handlers.go

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.getAllSnippets()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	// Create an instance of a templateData struct holding the slice of snippets.
	data := templateData{
		Snippets: snippets,
	}

	// Pass in the templateData struct when executing the template.
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	/* Extract the value of the id wildcard from the request using r.PathValue().
	   Convert it to an integer using the strconv.Atoi() function. */
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		log.Println(err)
	}

	// Use the SnippetModel's getSnippet() method to retrieve the data for a snippet.
	snippet, err := app.snippets.getSnippet(id)
	if err != nil {
		log.Println(err)
	}

	// Initialize a slice containing the paths to the base and view html files.
	files := []string{
		"./templates/base.html",
		"./templates/view.html",
	}

	// Parse the template files...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	// Create an instance of a templateData struct holding the snippet data.
	data := templateData{
		Snippet: snippet,
	}

	// Pass in the snippet data (a models.Snippet struct) as the final parameter.
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		files := []string{
			"./templates/base.html",
			"./templates/create.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
		}

		err = tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Print(err.Error())
		}
	case "POST":
		files := []string{
			"./templates/base.html",
			"./templates/create.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
		}

		msg := &Message{
			Title:   r.PostFormValue("title"),
			Content: r.PostFormValue("content"),
		}

		if !msg.Validate() {
			err = tmpl.ExecuteTemplate(w, "base", msg)
			if err != nil {
				log.Print(err.Error())
			}
		} else {
			title := r.FormValue("title")
			content := r.FormValue("content")

			err = app.snippets.insert(title, content)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func (app *application) editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

		// Use the SnippetModel's getSnippet() method to retrieve the data for a snippet.
		snippet, err := app.snippets.getSnippet(id)
		if err != nil {
			log.Println(err)
		}

		// Initialize a slice containing the paths to the base and view html files.
		files := []string{
			"./templates/base.html",
			"./templates/edit.html",
		}

		// Parse the template files...
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err)
		}

		msg := &Message{
			Title:   snippet.Title,
			Content: snippet.Content,
		}

		err = ts.ExecuteTemplate(w, "base", msg)
		if err != nil {
			log.Print(err.Error())
		}
	} else if r.Method == "POST" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

		files := []string{
			"./templates/base.html",
			"./templates/edit.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
		}

		msg := &Message{
			Title:   r.PostFormValue("title"),
			Content: r.PostFormValue("content"),
		}

		if !msg.Validate() {
			err = tmpl.ExecuteTemplate(w, "base", msg)
			if err != nil {
				log.Print(err.Error())
			}
		} else {
			title := r.FormValue("title")
			content := r.FormValue("content")

			err := app.snippets.update(id, title, content)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		log.Println(err)
	}

	err = app.snippets.delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./templates/base.html",
		"./templates/about.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
	}
}
