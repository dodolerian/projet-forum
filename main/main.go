package main

import (
	"database/sql"
	"fmt"
	"forum"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//forum.AddUsers(database, "personne1", "1234", "j'aime les jeux video", "test@gmail.com")
	//forum.ModifyBDD(database, 6, "petit test4")
	WebServer()
}

func WebServer() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/", Account)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 8871 : http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type accountStruct struct {
	UsernameError string
	PasswordError string
	MailError     string
}

func Account(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/account.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	defer database.Close()

	_, tmpUsername, _, _, tmpMail := forum.FetchUser(database)
	user := accountStruct{}

	// fmt.Println(r.FormValue("username"))
	// fmt.Println(r.FormValue("password"))
	// fmt.Println(r.FormValue("mail"))

	usernameForm := r.FormValue("username")
	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")

	if usernameForm != "" && passwordForm != "" && mailForm != "" {

		if ContainsStringArray(tmpUsername, usernameForm) {
			user = accountStruct{UsernameError: "nom déja utilisé"}

		} else if len(passwordForm) < 5 {
			user = accountStruct{PasswordError: "mot de passe trop court"}

		} else if ContainsStringArray(tmpMail, mailForm) {
			user = accountStruct{MailError: "adresse mail déja utilisé"}

		} else {
			forum.AddUsers(database, usernameForm, passwordForm, "", mailForm)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}

	err := tmpl.Execute(w, user)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))

	user := accountStruct{}

	err := tmpl.Execute(w, user)
	if err != nil {
		log.Fatal(err)
	}
}

func ContainsStringArray(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}
	return false
}
