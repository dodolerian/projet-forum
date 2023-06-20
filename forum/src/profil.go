<<<<<<< HEAD:forum/src/profil.go
package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ProfilPageStruct struct {
	Username          string
	ProfilDescription string
	Mail              string
}

func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/profil.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	description := r.FormValue("description")

	if r.Method == http.MethodPost {

		ContentPost := r.FormValue("ContentPost")
		tag := r.FormValue("tag")

		file, handler, err := r.FormFile("photo")
		if file != nil {
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			buff := make([]byte, handler.Size)

			_, err := file.Read(buff)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			filetype := http.DetectContentType(buff)

			if filetype != "image/jpeg" {
				fmt.Println("image type not good")
			} else {
				if handler.Size > 5000000 {
					fmt.Println("image to heavy")
				} else {
					AddPost(database, ContentPost, connectedUser[0], buff, tag)
					fmt.Println("post add image")

					currentXp, err := strconv.Atoi(connectedUser[5])
					currentId, err := strconv.Atoi(connectedUser[0])

					if err != nil {
						log.Fatal(err)
					}
					nextXp := currentXp + 10

					ModifyXpUser(database, currentId, nextXp)
				}
			}

		} else {
			AddPost(database, ContentPost, connectedUser[0], nil, tag)
			fmt.Println("post add without image")

			currentXp, err := strconv.Atoi(connectedUser[5])
			currentId, err := strconv.Atoi(connectedUser[0])

			if err != nil {
				log.Fatal(err)
			}
			nextXp := currentXp + 5

			ModifyXpUser(database, currentId, nextXp)
		}
	}

	// si le form est rempli alors change la valeur, empache qu'elle soit vide
	if description != "" {
		id, _ := strconv.Atoi(connectedUser[0])
		ModifyDescriptionUser(database, id, description)
	}

	// reprend l'utilisateur modifié a chaque fois
	idRefresh, username, password, profilDescription, mail, xp := FetchUserWithId(database, connectedUser[0])

	connectedUser = nil
	connectedUser = append(connectedUser, strconv.Itoa(idRefresh), username, password, profilDescription, mail, strconv.Itoa(xp))

	profilPage := ProfilPageStruct{
		Username:          connectedUser[1],
		ProfilDescription: connectedUser[3],
		Mail:              connectedUser[4],
	}
	err := tmpl.Execute(w, profilPage)
	if err != nil {
		log.Fatal(err)
	}
}
=======
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
		tag := r.FormValue("tag")

		file, handler, err := r.FormFile("photo")
		if file != nil {
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			buff := make([]byte, handler.Size)

			_, err := file.Read(buff)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			filetype := http.DetectContentType(buff)

			if filetype != "image/jpeg" {
				fmt.Println("image type not good")
			} else {
				if handler.Size > 5000000 {
					fmt.Println("image to heavy")
				} else {
					AddPost(database, ContentPost, connectedUser[0], buff, tag)
					fmt.Println("post add image")

					currentXp, err := strconv.Atoi(connectedUser[5])
					currentId, err := strconv.Atoi(connectedUser[0])

					if err != nil {
						log.Fatal(err)
					}
					nextXp := currentXp + 10

					ModifyXpUser(database, currentId, nextXp)
				}
			}

		} else {
			AddPost(database, ContentPost, connectedUser[0], nil, tag)
			fmt.Println("post add without image")

			currentXp, err := strconv.Atoi(connectedUser[5])
			currentId, err := strconv.Atoi(connectedUser[0])

			if err != nil {
				log.Fatal(err)
			}
			nextXp := currentXp + 5

			ModifyXpUser(database, currentId, nextXp)
		}
	}

	// si le form est rempli alors change la valeur, empache qu'elle soit vide
	if description != "" {
		id, _ := strconv.Atoi(connectedUser[0])
		ModifyDescriptionUser(database, id, description)
	}

	// reprend l'utilisateur modifié a chaque fois
	idRefresh, username, password, profilDescription, mail, xp := FetchUserWithId(database, connectedUser[0])

	connectedUser = nil
	connectedUser = append(connectedUser, strconv.Itoa(idRefresh), username, password, profilDescription, mail, strconv.Itoa(xp))

	xpInt, _ := strconv.Atoi(connectedUser[5])
	profilPage := ProfilPageStruct{
		ConnectedUserXp:   xpInt,
		Username:          connectedUser[1],
		ProfilDescription: connectedUser[3],
		Mail:              connectedUser[4],
	}
	err := tmpl.Execute(w, profilPage)
	if err != nil {
		log.Fatal(err)
	}
}
>>>>>>> 3b071879a73e125a9cbbab3d8374d2d25a1e592f:src/profil.go