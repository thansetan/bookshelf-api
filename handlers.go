package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateBookByID(w, r)
	case http.MethodDelete:
		deleteBookByID(w, r)
	default:
		getBookByID(w, r)
	}
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addBook(w, r)
	default:
		getBooks(w, r)
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var err error
	book := new(book)
	json.NewDecoder(r.Body).Decode(&book)
	if book.Name == "" {
		responseWriter(w, http.StatusBadRequest, "Gagal menambahkan buku. Mohon isi nama buku", nil)
		return
	}
	if book.ReadPage > book.PageCount {
		responseWriter(w, http.StatusBadRequest, "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount", nil)
		return
	}
	book.ID, err = gonanoid.New(16)
	if err != nil {
		responseWriter(w, http.StatusInternalServerError, "Gagal menambahkan buku", nil)
	}
	now := time.Now().UTC()
	book.InsertedAt = &now
	book.UpdatedAt = book.InsertedAt
	book.Finished = book.PageCount == book.ReadPage
	bookshelf = append(bookshelf, book)
	responseWriter(w, http.StatusCreated, "Buku berhasil ditambahkan", book)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var books []*bookSimpleResponse

	for _, book := range bookshelf {
		bookResp := &bookSimpleResponse{
			ID:        book.ID,
			Name:      book.Name,
			Publisher: book.Publisher,
		}
		if nameParam := params.Get("name"); nameParam != "" {
			if !strings.Contains(strings.ToLower(bookResp.Name), strings.ToLower(nameParam)) {
				continue
			}
		} else if finishedParam := params.Get("finished"); finishedParam != "" {
			if finishedParam == "1" && !book.Finished || finishedParam == "0" && book.Finished {
				continue
			}
		} else if readingParam := params.Get("reading"); readingParam != "" {
			if readingParam == "1" && !book.Reading || readingParam == "0" && book.Reading {
				continue
			}
		}
		books = append(books, bookResp)
	}
	responseWriter(w, http.StatusOK, "", map[string]any{"books": books})
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	var bookResp bookFullResponse
	for _, book := range bookshelf {
		if book.ID == id {
			bookResp = bookFullResponse(*book)
			responseWriter(w, http.StatusOK, "", map[string]any{"book": bookResp})
			return
		}
	}
	responseWriter(w, http.StatusNotFound, "Buku tidak ditemukan", nil)
}

func updateBookByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	var newBook book
	json.NewDecoder(r.Body).Decode(&newBook)
	if newBook.Name == "" {
		responseWriter(w, http.StatusBadRequest, "Gagal memperbarui buku. Mohon isi nama buku", nil)
		return
	}
	if newBook.ReadPage > newBook.PageCount {
		responseWriter(w, http.StatusBadRequest, "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount", nil)
		return
	}
	now := time.Now().UTC()
	newBook.UpdatedAt = &now
	for i, book := range bookshelf {
		if book.ID == id {
			newBook.InsertedAt = book.InsertedAt
			newBook.ID = book.ID
			bookshelf[i] = &newBook
			responseWriter(w, http.StatusOK, "Buku berhasil diperbarui", nil)
			return
		}
	}
	responseWriter(w, http.StatusNotFound, "Gagal memperbarui buku. Id tidak ditemukan", nil)
}

func deleteBookByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	for i, book := range bookshelf {
		if book.ID == id {
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			responseWriter(w, http.StatusOK, "Buku berhasil dihapus", nil)
			return
		}
	}
	responseWriter(w, http.StatusNotFound, "Buku gagal dihapus. Id tidak ditemukan", nil)
}
