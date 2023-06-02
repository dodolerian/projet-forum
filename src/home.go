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
	Post              []PostStruct
	NbrPost           int
}

type PostStruct struct {
	Id         int
	Author     int
	AuthorName string
	Content    string
	Like       bool
	Dislike    bool
	Date       string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}

	RecuperationLike()

	allPost = nil
	allPost := recuperationPost()
	allPostFinal := []PostStruct{}

	connectedUserId, _ := strconv.Atoi(connectedUser[0])

	if r.Method == http.MethodPost {
		ContentComment := r.FormValue("ContentComment")
		AddComment(database, ContentComment, connectedUser[0], "1")
	}

	if r.Method == http.MethodPost {
		likeIdPostStr := r.FormValue("like")
		likeIdPost, _ := strconv.Atoi(likeIdPostStr)
		isLiked := LikeOnPost(connectedUserId, likeIdPost, allLikeList)
		if isLiked == true {
			DeleteLike(database, connectedUserId, likeIdPost)
		} else {
			AddLike(database, connectedUserId, likeIdPost)
		}
	}

	//LIKE

	RecuperationLike()

	for i := 0; i < len(allPost); i++ {
		_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		postFinalIntoStruc := PostStruct{
			Id:         allPost[i].Id,
			Author:     allPost[i].Author,
			AuthorName: username,
			Content:    allPost[i].Content,
			Like:       isLiked,
			Dislike:    true,
			Date:       allPost[i].Date,
		}

		allPostFinal = append(allPostFinal, postFinalIntoStruc)
	}

	if len(connectedUser) > 0 {
		homePage = HomePageStruct{
			Post:    allPostFinal,
			NbrPost: len(allPost),
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
