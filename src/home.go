package forum

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HomePageStruct struct {
	Post        []PostStruct
	NbrPost     int
	Comments    []recuperationCommentFromDb
	User        []UserStruct
	IsConnected bool
}

type CommentStruct struct {
	IdPost     int
	IdAuthor   int
	AuthorName string
	Content    string
	Like       int
	Dislike    int
	Date       string
	Image      string
	Image      string
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
	Image       string
	IsImage     bool
	IsConnected bool
}

type UserStruct struct {
	Id                int
	Username          string
	ProfilDescription string
}

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}
	// invité := "invité"
	// invité := "invité"
	RecuperationLike()
	RecuperationDislike()

	allPost = nil
	allPost := recuperationPost()
	allPostFinal := []PostStruct{}

	/* Recuperation User */

	allUsers = nil
	allUsers := recuperationUser()
	User := []UserStruct{}

	//.....

	if len(connectedUser) == 0 {
		connectedUser = append(connectedUser, "-1")
	}

	connectedUserId, _ := strconv.Atoi(connectedUser[0])

	/* Author Post */
	if r.Method == http.MethodPost {
		IdAuthor := r.FormValue("author")
		for i := 0; i < len(allUsers); i++ {
			if strconv.Itoa(allUsers[i].Id) == IdAuthor {
				name := allUsers[i].Username
				description := allUsers[i].ProfilDescription
				if description == "" {
					description = "Pas de description"
				}
				userStruct := UserStruct{}
				userStruct = UserStruct{
					Id:                allUsers[i].Id,
					Username:          name,
					ProfilDescription: description,
				}
				User = append(User, userStruct)
			}
		}
	}

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

	for i := len(allPost) - 1; i >= 0; i-- {
		_, username, _, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		isDisliked := DislikeOnPost(connectedUserId, allPost[i].Id, allDislikeList)
		_, username, _, _, _, _ = FetchUserWithId(database, strconv.Itoa(allPost[i].Author))
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

		imgBase64Str := base64.StdEncoding.EncodeToString(allPost[i].Image)
		isImage := true
		if imgBase64Str == "" {
			isImage = false
		}

		postFinalIntoStruc := PostStruct{
			Id:          allPost[i].Id,
			Author:      allPost[i].Author,
			AuthorName:  username,
			Content:     content,
			Like:        isLiked,
			Dislike:     isDisliked,
			Date:        allPost[i].Date,
			Comments:    allCommentOfThisPost,
			Image:       imgBase64Str,
			IsImage:     isImage,
			IsConnected: true,
		}

		if len(connectedUser) == 1 {
			postFinalIntoStruc.IsConnected = false
		}
		/* Add comments of this post */
		for j := 0; j < len(allComment); j++ {
			if allPost[i].Id == allComment[j].IdPost {
				_, username, _, _, _, _ := FetchUserWithId(database, strconv.Itoa(allComment[j].IdAuthor))
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
			Post:        allPostFinal,
			NbrPost:     len(allPost),
			Comments:    allComment,
			User:        User,
			IsConnected: true,
		}

	} else {
		homePage = HomePageStruct{
			// IdAuthor:          "-1",
			// Username:          invité,
			// ProfilDescription: invité,
			// Mail:              invité,

			Post:        allPostFinal,
			NbrPost:     len(allPost),
			Comments:    allComment,
			User:        User,
			IsConnected: false,
		}
	}

	fmt.Println(connectedUser)
	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
