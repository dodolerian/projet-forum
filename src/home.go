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
	Comments          []recuperationCommentFromDb
}

type CommentStruct struct {
	IdPost     int
	IdAuthor   int
	AuthorName string
	Content    string
	Like       int
	Dislike    int
	Date       string
}

type PostStruct struct {
	Id         int
	Author     int
	AuthorName string
	Content    string
	Like       bool
	Dislike    bool
	Date       string
	Comments   []CommentStruct
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
		IdPost := r.FormValue("idPost")
		ContentComment := r.FormValue("ContentComment")
		AddComment(database, ContentComment, connectedUser[0], IdPost)
	}
	allPost = nil
	allPost := recuperationPost()

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

	allComment = nil
	allComment := recuperationComment()

	allCommentOfThisPost := []CommentStruct{}

	allPostFinal := []PostStruct{}
	RecuperationLike()
	RecuperationDislike()

	for i := 0; i < len(allPost); i++ {
		_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		isDisliked := DislikeOnPost(connectedUserId, allPost[i].Id, allDislikeList)
		postFinalIntoStruc := PostStruct{
			Id:         allPost[i].Id,
			Author:     allPost[i].Author,
			AuthorName: username,
			Content:    allPost[i].Content,
			Like:       isLiked,
			Dislike:    isDisliked,
			Date:       allPost[i].Date,
			Comments:   allCommentOfThisPost,
		}
		/* Add comments of this post */
		for j := 0; j < len(allComment); j++ {
			if allPost[i].Id == allComment[j].IdPost {
				_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allComment[j].IdAuthor))
				commentIntoStruc := CommentStruct{
					IdPost:     allComment[j].IdPost,
					IdAuthor:   allComment[j].IdAuthor,
					AuthorName: username,
					Content:    allComment[j].Content,
					Like:       allComment[j].Like,
					Dislike:    allComment[j].Dislike,
					Date:       allComment[j].Date,
				}
				postFinalIntoStruc.Comments = append(postFinalIntoStruc.Comments, commentIntoStruc)
			}
		}

		allPostFinal = append(allPostFinal, postFinalIntoStruc)
	}

	if len(connectedUser) > 0 {
		homePage = HomePageStruct{
			IdAuthor:          connectedUser[0],
			Username:          connectedUser[1],
			ProfilDescription: connectedUser[3],
			Mail:              connectedUser[4],
			ContentPost:       allPost[0].Content,
			LikePost:          allPost[0].Like,
			DyslikePost:       allPost[0].Dislike,
			DatePost:          allPost[0].Date,
			Post:              allPostFinal,
			NbrPost:           len(allPost),
			ContentComment:    allComment[0].Content,
			IdPostComment:     allComment[0].IdPost,
			LikeComment:       allComment[0].Like,
			DyslikeComment:    allComment[0].Dislike,
			Comments:          allComment,
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
