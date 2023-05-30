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
	ContentComment    string
	AuthorComment     string
	IdPostComment     int
	LikeComment       int
	DyslikeComment    int
	Post              []recuperationPostFromDb
	NbrPost           int
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}

	if r.Method == http.MethodPost {
		ContentComment := r.FormValue("ContentComment")
		AddComment(database, ContentComment, connectedUser[0], "1")
	}

	allPost = nil
	allPost := recuperationPost()

	allComment = nil
	allComment = recuperationComment()

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
			// ContentComment:    allPost[0].ContentComment,
			// AuthorComment:     allPost[0].AuthorComment,
			// IdPostComment:     allPost[0].IdPostComment,
			// LikeComment:       allPost[0].LikeComment,
			// DyslikeComment:    allPost[0].DislikeComment,
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
