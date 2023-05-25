package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var connectedUser []string

func WebServer() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/create", CreateAccount)
	http.HandleFunc("/", ConnexionAccount)
	http.HandleFunc("/profil", Profil)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 8871 : http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type createAccountStruct struct {
	UsernameError string
	PasswordError string
	MailError     string
}

type HomePageStruct struct {
	IdAuthor          string
	Username          string
	ProfilDescription string
	Mail              string
	ContentPost       string
	AuthorPost        string
	LikePost          int
	DyslikePost       int
	DatePost          string
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/createAccount.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	defer database.Close()

	_, tmpUsername, _, _, tmpMail := FetchAllUser(database)
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
			//hash
			AddUsers(database, usernameForm, passwordForm, "", mailForm)
			id, username, password, profilDescription, mail := FetchUserWithName(database, usernameForm)
			connectedUser = append(connectedUser, strconv.Itoa(id), username, password, profilDescription, mail)
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

	defer database.Close()

	_, _, _, _, tmpMail := FetchAllUser(database)
	accountPage := createAccountStruct{}

	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")

	if passwordForm != "" && mailForm != "" {
		if !ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail pas trouvé"}
		} else {
			id, username, password, profilDescription, mail := FetchUserWithMail(database, mailForm)

			//dehash
			if passwordForm != password {
				accountPage = createAccountStruct{PasswordError: "mot de passe faux"}
			} else {
				connectedUser = append(connectedUser, strconv.Itoa(id), username, password, profilDescription, mail)
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}
	}

	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}

	if r.Method == http.MethodPost {
		ContentPost := r.FormValue("ContentPost")
		AddPost(database, ContentPost, homePage)
	}
	allPost := recuperationPost()

	_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[2].Author))

	if len(connectedUser) > 0 {
		homePage = HomePageStruct{
			IdAuthor:          connectedUser[0],
			Username:          connectedUser[1],
			ProfilDescription: connectedUser[3],
			Mail:              connectedUser[4],
			ContentPost:       allPost[2].Content,
			AuthorPost:        username,
			LikePost:          allPost[2].Like,
			DyslikePost:       allPost[2].Dislike,
			DatePost:          allPost[2].Date,
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/profil.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()

	description := r.FormValue("description")

	// si le form est rempli alors change la valeur, empache qu'elle soit vide
	if description != "" {
		id, _ := strconv.Atoi(connectedUser[0])
		ModifyDescriptionUser(database, id, description)
	}

	// reprend l'utilisateur modifié a chaque fois
	idRefresh, username, password, profilDescription, mail := FetchUserWithId(database, connectedUser[0])
	connectedUser = nil
	connectedUser = append(connectedUser, strconv.Itoa(idRefresh), username, password, profilDescription, mail)

	profilPage := HomePageStruct{
		Username:          connectedUser[1],
		ProfilDescription: connectedUser[3],
		Mail:              connectedUser[4],
	}

	err := tmpl.Execute(w, profilPage)
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
