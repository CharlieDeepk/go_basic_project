package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"

	"github.com/CharlieDeepk/go_basic_project/pkg/models"
	"github.com/CharlieDeepk/go_basic_project/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, encodeErr := json.Marshal(newBooks)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id, parseErr := strconv.ParseInt(params["bookId"], 10, 64)
	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while Parsing")
		return
	}

	book, _ := models.GetBookById(Id)

	res, encodeErr := json.Marshal(book)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	incomingBook := &models.Book{}
	utils.ParseBody(r, incomingBook)
	b := incomingBook.CreateBook()

	res, encodeErr := json.Marshal(b)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, parseErr := strconv.ParseInt(params["bookId"], 10, 64)
	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while parsing")
		return
	}

	b := models.DeleteBook(Id)

	res, encodeErr := json.Marshal(b)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook_v1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, parseErr := strconv.ParseInt(params["bookId"], 10, 64)
	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while parsing")
		return
	}
	//deletion
	models.DeleteBook(Id)
	//creation
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	updatedBook.ID = uint(Id)
	b := updatedBook.CreateBook()
	res, encodeErr := json.Marshal(b)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook_v2(w http.ResponseWriter, r *http.Request) {
	//separating out Id
	params := mux.Vars(r)
	ID, parseErr := strconv.ParseInt(params["bookId"], 10, 64)
	if parseErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while parsing")
		return
	}

	//incoming body data
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	//updating row
	existingBook, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		existingBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		existingBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		existingBook.Publication = updatedBook.Publication
	}

	db.Save(&existingBook)

	//
	res, encodeErr := json.Marshal(existingBook)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error while encoding")
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
