package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/davoodmossgreen/books-i/internal/config"
	"github.com/davoodmossgreen/books-i/internal/models"
	"github.com/davoodmossgreen/books-i/internal/utils"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"

)

var db *gorm.DB

func init(){
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&models.User{})
}

var user models.User
	
var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")


	var sqluser models.User

	db.Where("email = ?", user.Email).First(&sqluser)
	if sqluser.Email != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email already exists"))
		return
	}

	db.Where("username = ?", user.Username).First(&sqluser)
	if sqluser.Username != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username already exists"))
		return
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatalln("error while hashing password")
	}

	user.Password = password

	if user.Username == "admin" {
		user.User_type = "admin"
	} else {
		user.User_type = "user"
	}


	db.NewRecord(user)
	db.Create(&user)

	session, _ := store.Get(r, "session")
	session.Values["username"] = user.Username
	session.Save(r, w)
	redirectTarget := "/"
	http.Redirect(w, r, redirectTarget, http.StatusFound)
	} else {
		utils.TplParse(w, "signup.html", nil)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
	redirectTarget := "/"
	password := r.FormValue("password")
	email := r.FormValue("email")

	var authInfo models.User
	
	if email == "" {
		msg := "please provide email"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}

	if password == "" {
		msg := "please provide password"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}


	db.Where("email = ?", email).First(&authInfo)
	if authInfo.Email == "" {
		msg := "email or password is incrorrect"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}

	if !utils.CheckPasswordHash(password, authInfo.Password) {
		msg := "email or password is incrorrecttt"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["username"] = authInfo.Username
	session.Save(r, w)
	redirectTarget = "/mybooks"
	http.Redirect(w, r, redirectTarget, http.StatusFound)
	} else {
	utils.TplParse(w, "login.html", nil)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
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
		searchVal := r.FormValue("username")
		var searchUser []models.User
		db.Raw("SELECT username, email, user_type, created_at FROM users WHERE username=?", searchVal).Scan(&searchUser)
		type Upl struct {
			Ok bool
			Users []models.User
		}
	
		var uppl Upl
		uppl.Ok = ok
		uppl.Users = searchUser

		utils.TplParse(w, "found.html", uppl)
	} else {
		if value != "admin" {
		msg := "sorry ... this feature is available only for admins"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return
		}
		utils.TplParse(w, "search.html", ok)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	val, ok := session.Values["username"]
	fmt.Println("ok:", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	value := val.(string)

	if r.Method == "POST" {
		yes := r.FormValue("Yes")
		no := r.FormValue("No")

		if no == "no" {
			http.Redirect(w, r, "/", http.StatusFound)
		} else if yes == "yes" {
			db.Exec("DELETE FROM users WHERE username=?", value)
			clearSession(w)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		utils.TplParse(w, "deleteAccount.html", ok)
	}
}