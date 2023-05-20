package main

import (
	"encoding/json" // To work with Json
	"log"           //-log error
	"net/http"      //-to work with http like API and stuffs
	//"math/rand"//-whe we create create a book want to create a id for random numbers.
	// "github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("Book not found")
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books[index] = Book{
				ID:     params["id"],
				Title:  book.Title,
				Author: book.Author,
			}
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode("Book not found")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode("Book not found")
}

func main() {
	//An HTTP router is an application that ties the requested URL to the processing of the response.
	router := mux.NewRouter() //Initializating Router

	books = append(books, Book{ID: "1", Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: "2", Title: "Book 2", Author: "Author 2"})

	//Route handlers which will establish our endpoints for our API.
	//Route handlers are the blocks of code that handle logic for your routes.
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
