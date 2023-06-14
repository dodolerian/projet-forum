package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func ConnexionAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/connexionAccount.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	if r.Method == http.MethodPost {
		r.FormValue("deconnection")
		connectedUser = nil

	}

	defer database.Close()

	_, _, _, _, tmpMail := FetchAllUser(database)
	accountPage := createAccountStruct{}

	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")
	rememberMe := r.FormValue("rememberMe")
	if rememberMe != "" {

		fmt.Println(("on est ici"))
	}
	if passwordForm != "" && mailForm != "" {
		if !ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail pas trouvé"}
		} else {
			id, username, hashpass, profilDescription, mail := FetchUserWithMail(database, mailForm)
			//dehash
			if !CheckPasswordHash(passwordForm, hashpass) {
				accountPage = createAccountStruct{PasswordError: "mot de passe faux"}
			} else {
				connectedUser = nil
				connectedUser = append(connectedUser, strconv.Itoa(id), username, hashpass, profilDescription, mail)
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}
	}

	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}
