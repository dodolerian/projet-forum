package forum

import (
	"database/sql"
	"encoding/base64"
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
	Image      string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}

	RecuperationLike()
	RecuperationDislike()

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
		dislikeIdPostStr := r.FormValue("dislike")

		likeIdPost, _ := strconv.Atoi(likeIdPostStr)
		dislikeIdPost, _ := strconv.Atoi(dislikeIdPostStr)

		if likeIdPostStr == "" {
			isLiked := LikeOnPost(connectedUserId, dislikeIdPost, allLikeList)
			isDisliked := DislikeOnPost(connectedUserId, dislikeIdPost, allDislikeList)

			if isLiked == true {
				DeleteLike(database, connectedUserId, dislikeIdPost)
			}
			if isDisliked == true {
				DeleteDislike(database, connectedUserId, dislikeIdPost)
			} else {
				AddDislike(database, connectedUserId, dislikeIdPost)
			}
		}

		if dislikeIdPostStr == "" {
			isLiked := LikeOnPost(connectedUserId, likeIdPost, allLikeList)
			isDisliked := DislikeOnPost(connectedUserId, likeIdPost, allDislikeList)

			if isDisliked == true {
				DeleteDislike(database, connectedUserId, likeIdPost)
			}
			if isLiked == true {
				DeleteLike(database, connectedUserId, likeIdPost)
			} else {
				AddLike(database, connectedUserId, likeIdPost)
			}
		}

	}

	//LIKE

	RecuperationLike()
	RecuperationDislike()

	for i := len(allPost) - 1; i >= 0; i-- {
		_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		isDisliked := DislikeOnPost(connectedUserId, allPost[i].Id, allDislikeList)

		imgBase64Str := base64.StdEncoding.EncodeToString(allPost[i].Image)

		postFinalIntoStruc := PostStruct{
			Id:         allPost[i].Id,
			Author:     allPost[i].Author,
			AuthorName: username,
			Content:    allPost[i].Content,
			Like:       isLiked,
			Dislike:    isDisliked,
			Date:       allPost[i].Date,
			Image:      imgBase64Str,
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
