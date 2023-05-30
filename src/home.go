package forum

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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
	Post              []recuperationPostFromDb
	NbrPost           int
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}

	if r.Method == http.MethodPost {
		ContentPost := r.FormValue("ContentPost")
		AddPost(database, ContentPost, connectedUser[0])
	}

	allPost = nil
	allPost := recuperationPost()

	_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[0].Author))

	if len(connectedUser) > 0 {
		homePage = HomePageStruct{
			IdAuthor:          connectedUser[0],
			Username:          connectedUser[1],
			ProfilDescription: connectedUser[3],
			Mail:              connectedUser[4],
			ContentPost:       allPost[0].Content,
			AuthorPost:        username,
			LikePost:          allPost[0].Like,
			DyslikePost:       allPost[0].Dislike,
			DatePost:          allPost[0].Date,
			Post:              allPost,
			NbrPost:           len(allPost),
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
