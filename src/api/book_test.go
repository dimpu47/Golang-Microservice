package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{
		Title:  "Go Cloud Book",
		Author: "Gaurav Choudhary",
		ISBN:   "013456789",
	}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Go Cloud Book","Author":"Gaurav Choudhary","ISBN":"013456789"}`, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Go Cloud Book","Author":"Gaurav Choudhary","isbn":"013456789"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Go Cloud Book", Author: "Gaurav Choudhary", ISBN: "013456789"}, book, "Book JSON unmarshalling wrong.")
}
