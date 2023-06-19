package forum

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type createAccountStruct struct {
	UsernameError string
	PasswordError string
	MailError     string
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/createAccount.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	defer database.Close()

	_, tmpUsername, _, _, tmpMail, _ := FetchAllUser(database)
	accountPage := createAccountStruct{}

	usernameForm := r.FormValue("username")
	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")

	if usernameForm != "" && passwordForm != "" && mailForm != "" {

		if ContainsStringArray(tmpUsername, usernameForm) {
			accountPage = createAccountStruct{UsernameError: "nom déja utilisé"}

		} else if len(passwordForm) < 5 {
			accountPage = createAccountStruct{PasswordError: "mot de passe trop court"}

		} else if ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail déja utilisé"}

		} else {
			hashpass, _ := HashPassword(passwordForm)


			AddUsers(database, usernameForm, hashpass, "", mailForm)
			id, username, password, profilDescription, mail, xp := FetchUserWithName(database, usernameForm)
			connectedUser = nil
			connectedUser = append(connectedUser, strconv.Itoa(id), username, password, profilDescription, mail, strconv.Itoa(xp))
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}
