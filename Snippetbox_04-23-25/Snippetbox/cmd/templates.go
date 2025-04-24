// templates.go

package main

/* Define a templateData type to act as the holding structure for
   any dynamic data that we want to pass to our HTML templates. */

type templateData struct {
	Snippets []Snippet
	Snippet Snippet
}
