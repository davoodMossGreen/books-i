package controllers

import (
	"fmt"
	"net/http"

	"github.com/davoodmossgreen/books-i/internal/config"
	"github.com/davoodmossgreen/books-i/internal/models"
	"github.com/davoodmossgreen/books-i/internal/utils"
)


func init(){
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&models.Book{})
}

func init(){
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&models.Note{})
}

var book models.Book


func Index(w http.ResponseWriter, r *http.Request) {
	
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]

	utils.TplParse(w, "index.html", ok)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method == "POST" {
	val := session.Values["username"]
	value := val.(string)

	var newbook models.Book

	newbook.Uploader = value
	newbook.Author = r.FormValue("author")
	newbook.Name = r.FormValue("title")
	newbook.Publication = r.FormValue("publication")
	newbook.Description = r.FormValue("description")

	db.NewRecord(newbook)
	db.Create(&newbook)

	redirectTarget := "/mybooks"
	http.Redirect(w, r, redirectTarget, http.StatusFound)
} else {
	utils.TplParse(w, "add.html", ok)
}

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	
	val := session.Values["username"]
	value := val.(string)
	var allBooks []models.Book
	db.Where("uploader=?", value).Order("author").Find(&allBooks)
	
	type Upl struct {
		Ok bool
		Books []models.Book
	}

	var uppl Upl

	uppl.Ok = ok
	uppl.Books = allBooks

	utils.TplParse(w, "mybooks.html", uppl)
}



func DeleteBook(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	val := session.Values["username"]
	value := val.(string)
	
	if r.Method == "POST" {
	bookName := r.FormValue("name")
	db.Where("name = ?", bookName).Delete(book)
	http.Redirect(w, r, "/mybooks", http.StatusFound)
	} else {
		type Upl struct {
			Ok bool
			Books []models.Book
		}

		var allBooks []models.Book
		db.Where("uploader=?", value).Order("created_at desc").Find(&allBooks)

		var uppl Upl
		uppl.Ok = ok
		uppl.Books = allBooks

		utils.TplParse(w, "deletebook.html", uppl)
	}
}

func Notes(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	val, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	value := val.(string)
	var note models.Note
	

	if r.Method == "POST" {
		note.Heading = r.FormValue("heading")
		note.Note = r.FormValue("note")
		note.Username = value

		db.NewRecord(note)
		db.Create(&note)
		http.Redirect(w, r, "/notes", http.StatusFound)
	} else {
		var allnotes []models.Note
		db.Raw("SELECT heading, note, username, created_at FROM notes WHERE username=?", value).Scan(&allnotes)
		type Upll struct {
			Ok bool
			Notes []models.Note
		}
		var notes Upll
		notes.Ok = ok
		notes.Notes = allnotes

		utils.TplParse(w, "notes.html", notes)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	val, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	value := val.(string)
	if r.Method == "POST" {
	note := r.FormValue("heading")
	db.Exec("DELETE FROM notes WHERE username=? AND heading=?", value, note)
	http.Redirect(w, r, "/notes", http.StatusFound)
	}
}

func History(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	val, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	value := val.(string)
	var books []models.Book
	db.Raw("SELECT author, name, publication, created_at FROM books WHERE uploader=? ORDER BY created_at", value).Scan(&books)
	type Upl struct {
		Ok bool
		Books []models.Book
	}
	var uppl Upl
	uppl.Ok = ok
	uppl.Books = books

	utils.TplParse(w, "history.html", uppl)
}