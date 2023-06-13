package forum

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	IsConnected       bool
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
	Id          int
	Author      int
	AuthorName  string
	Content     string
	Like        bool
	Dislike     bool
	Date        string
	Comments    []CommentStruct
	IsConnected bool
}

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}
	invité := "invité"
	RecuperationLike()
	RecuperationDislike()

	allPost = nil
	allPost := recuperationPost()
	allPostFinal := []PostStruct{}
	if len(connectedUser) == 0 {
		connectedUser = append(connectedUser, "-1")
	}
	connectedUserId, _ := strconv.Atoi(connectedUser[0])

	/* COMMENTS */
	if r.Method == http.MethodPost {
		IdPost := r.FormValue("idPost")
		ContentComment := r.FormValue("ContentComment")
		AddComment(database, ContentComment, connectedUser[0], IdPost)
	}

	/* LIKE */
	if r.Method == http.MethodPost {
		likeIdPostStr := r.FormValue("like")
		dislikeIdPostStr := r.FormValue("dislike")

		likeIdPost, _ := strconv.Atoi(likeIdPostStr)
		dislikeIdPost, _ := strconv.Atoi(dislikeIdPostStr)

		if likeIdPostStr == "" {
			isLiked := LikeOnPost(connectedUserId, dislikeIdPost, allLikeList)
			isDisliked := DislikeOnPost(connectedUserId, dislikeIdPost, allDislikeList)

			if isLiked {
				DeleteLike(database, connectedUserId, dislikeIdPost)
			}
			if isDisliked {
				DeleteDislike(database, connectedUserId, dislikeIdPost)
			} else {
				AddDislike(database, connectedUserId, dislikeIdPost)
			}
		}

		if dislikeIdPostStr == "" {
			isLiked := LikeOnPost(connectedUserId, likeIdPost, allLikeList)
			isDisliked := DislikeOnPost(connectedUserId, likeIdPost, allDislikeList)

			if isDisliked {
				DeleteDislike(database, connectedUserId, likeIdPost)
			}
			if isLiked {
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

	allPostFinal = []PostStruct{}
	RecuperationLike()
	RecuperationDislike()

	for i := 0; i < len(allPost); i++ {
		_, username, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		isDisliked := DislikeOnPost(connectedUserId, allPost[i].Id, allDislikeList)
		_, username, _, _, _ = FetchUserWithId(database, strconv.Itoa(allPost[i].Author))
		/* Check valide post */
		checkPost := strings.Split(allPost[i].Content, "")
		limite := 0

		for v := 0; v < len(checkPost); v++ {
			if limite >= 27 && checkPost[v] != " " {
				checkPost[v] = " "
				limite = 0
			}
			limite++
		}

		content := strings.Join(checkPost, "")
		postFinalIntoStruc := PostStruct{
			Id:          allPost[i].Id,
			Author:      allPost[i].Author,
			AuthorName:  username,
			Content:     content,
			Like:        isLiked,
			Dislike:     isDisliked,
			Date:        allPost[i].Date,
			Comments:    allCommentOfThisPost,
			IsConnected: true,
		}

		if len(connectedUser) == 1 {
			postFinalIntoStruc.IsConnected = false
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
	if len(connectedUser) > 1 {
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
			IsConnected:       true,
		}

	} else {
		homePage = HomePageStruct{

			IdAuthor:          "-1",
			Username:          invité,
			ProfilDescription: invité,
			Mail:              invité,
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
			IsConnected:       false,
		}
	}
	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
