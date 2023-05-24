package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

var tpl *template.Template

func TplParse(w http.ResponseWriter, html string, data interface{}) {
	tpl, _ = tpl.ParseGlob("../templates/*.html")
	err := tpl.ExecuteTemplate(w, html, data)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

