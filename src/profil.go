package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)


func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/profil.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	description := r.FormValue("description")
	errorPost :=""

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
				errorPost = "type de l'image pas bon"
			} else {
				if handler.Size > 5000000 {
					fmt.Println("image to heavy")
					errorPost = "image trop grande"
				} else {
					if !VerifyPostContent(ContentPost) {
						AddPost(database, ContentPost, connectedUser[0], buff, tag)
						fmt.Println("post add image")

						currentXp, err := strconv.Atoi(connectedUser[5])
						currentId, err := strconv.Atoi(connectedUser[0])

						if err != nil {
							log.Fatal(err)
						}
						nextXp := currentXp + 10

						ModifyXpUser(database, currentId, nextXp)
					} else {
						errorPost = "un mot banni a été utilisé"
					}
				}
			}

		} else {
			if !VerifyPostContent(ContentPost) {
				AddPost(database, ContentPost, connectedUser[0], nil, tag)
				fmt.Println("post add without image")

				currentXp, err := strconv.Atoi(connectedUser[5])
				currentId, err := strconv.Atoi(connectedUser[0])

				if err != nil {
					log.Fatal(err)
				}
				nextXp := currentXp + 5

				ModifyXpUser(database, currentId, nextXp)
			} else {
				fmt.Println(VerifyPostContent(ContentPost))
				
				errorPost = "un mot banni a été utilisé"
			}
		}
	}

	// si le form est rempli alors change la valeur, empeche qu'elle soit vide
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

	if errorPost != "" {
		profilPage = ProfilPageStruct{
			ConnectedUserXp:   xpInt,
			Username:          connectedUser[1],
			ProfilDescription: connectedUser[3],
			Mail:              connectedUser[4],
			Error: errorPost,
		}
	}
	
	err := tmpl.Execute(w, profilPage)
	if err != nil {
		log.Fatal(err)
	}
}

func VerifyPostContent(content string) bool {
	bannedWords := []string{"nazi","gay","lgbt","suicide","terrorisme", "fuck", "connard", "pute", "salope", "enculé"}
	for i:=0;i<len(bannedWords);i++{
		if strings.Contains(content,bannedWords[i]) {
			return true
		}	
	}
	return false
}