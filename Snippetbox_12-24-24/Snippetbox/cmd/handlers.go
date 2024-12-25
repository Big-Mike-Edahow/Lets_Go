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
		"./templates/nav.html",
		"./templates/index.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err)
	}

	err = tmpl.ExecuteTemplate(w, "base", snippets)
	if err != nil {
		log.Print(err.Error())
	}
}

func (app *application) viewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

	snippet, err := app.snippets.getOneSnippet(id)
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/nav.html",
		"./templates/view.html",
	}

	// Parse the template files...
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		files := []string{
			"./templates/base.html",
			"./templates/nav.html",
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
			"./templates/nav.html",
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

			_, err = app.snippets.insert(title, content)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func (app *application) editHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

		log.Println("Get request. Id is:", id)

		files := []string{
			"./templates/base.html",
			"./templates/nav.html",
			"./templates/edit.html",
		}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
		}
		snippet, err := app.snippets.getOneSnippet(id)
		if err != nil {
			log.Println(err)
		}

		msg := &Message{
			Title:   snippet.Title,
			Content: snippet.Content,
		}

		err = tmpl.ExecuteTemplate(w, "base", msg)
		if err != nil {
			log.Print(err.Error())
		}
	case "POST":
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

		log.Println("Post request. Id is:", id)

		files := []string{
			"./templates/base.html",
			"./templates/nav.html",
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
