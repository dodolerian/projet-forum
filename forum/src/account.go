package forum

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func ContainsStringArray(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}
	return false
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

func ConnexionAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/connexionAccount.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	if r.Method == http.MethodPost {
		r.FormValue("deconnection")
		connectedUser = nil
		connectedUser = nil

	}

	defer database.Close()

	_, _, _, _, tmpMail, _ := FetchAllUser(database)
	accountPage := createAccountStruct{}

	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")
	if passwordForm != "" && mailForm != "" {
		if !ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail pas trouvé"}
		} else {
			id, username, hashpass, profilDescription, mail, xp := FetchUserWithMail(database, mailForm)
			//dehash
			if !CheckPasswordHash(passwordForm, hashpass) {
				accountPage = createAccountStruct{PasswordError: "mot de passe faux"}
			} else {
				connectedUser = nil
				connectedUser = append(connectedUser, strconv.Itoa(id), username, hashpass, profilDescription, mail, strconv.Itoa(xp))
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}
	}

	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}
