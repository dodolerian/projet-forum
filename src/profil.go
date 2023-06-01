package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/profil.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()

	description := r.FormValue("description")

	if r.Method == http.MethodPost {
		ContentPost := r.FormValue("ContentPost")
		AddPost(database, ContentPost, connectedUser[0])
	}

	// si le form est rempli alors change la valeur, empache qu'elle soit vide
	if description != "" {
		id, _ := strconv.Atoi(connectedUser[0])
		ModifyDescriptionUser(database, id, description)
	}

	// reprend l'utilisateur modifié a chaque fois
	idRefresh, username, password, profilDescription, mail := FetchUserWithId(database, connectedUser[0])
		fmt.Println("problema la")

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