package forum

import (
	"database/sql"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	homePage := HomePageStruct{}
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

	tag := ""

	/* COMMENTS */
	if r.Method == http.MethodPost {
		IdPost := r.FormValue("idPost")
		tag = r.FormValue("tag")
		ContentComment := r.FormValue("ContentComment")
		AddComment(database, ContentComment, connectedUser[0], IdPost)
	}

	/* LIKE */
	if r.Method == http.MethodPost {
		likeIdPostStr := r.FormValue("like")
		dislikeIdPostStr := r.FormValue("dislike")

		UseLike(database, likeIdPostStr, dislikeIdPostStr, connectedUserId)
	}

	if r.Method == http.MethodPost {
		likeIdCommentStr := r.FormValue("likeComment")
		dislikeIdCommentStr := r.FormValue("dislikeComment")

		UseLikeComment(database, likeIdCommentStr, dislikeIdCommentStr, connectedUserId)
	}

	allComment = nil
	allComment := recuperationComment()

	allCommentOfThisPost := []CommentStruct{}

	allPostFinal = nil
	allPostFinal = []PostStruct{}

	RecuperationLike()
	RecuperationDislike()
	RecuperationLikeComment()
	RecuperationDislikeComment()

	for i := len(allPost) - 1; i >= 0; i-- {
		_, username, _, _, _, _ := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))

		isLiked := LikeOnPost(connectedUserId, allPost[i].Id, allLikeList)
		isDisliked := DislikeOnPost(connectedUserId, allPost[i].Id, allDislikeList)
		_, username, _, _, _, xpIntPost := FetchUserWithId(database, strconv.Itoa(allPost[i].Author))
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
			Tag:         allPost[i].Tag,
			Xp:          xpIntPost,
		}

		if len(connectedUser) == 1 {
			postFinalIntoStruc.IsConnected = false
		}
		/* Add comments of this post */
		for j := 0; j < len(allComment); j++ {
			if allPost[i].Id == allComment[j].IdPost {
				isLikedComment := LikeOnComment(connectedUserId, allComment[j].IdComment, allLikeCommentList)
				isDislikedComment := DislikeOnComment(connectedUserId, allComment[j].IdComment, allDislikeCommentList)
				_, username, _, _, _, _ := FetchUserWithId(database, strconv.Itoa(allComment[j].IdAuthor))
				commentIntoStruc := CommentStruct{
					IdPost:     allComment[j].IdPost,
					IdAuthor:   allComment[j].IdAuthor,
					AuthorName: username,
					Content:    allComment[j].Content,
					Like:       isLikedComment,
					Dislike:    isDislikedComment,
					IdComment:  allComment[j].IdComment,
				}
				postFinalIntoStruc.Comments = append(postFinalIntoStruc.Comments, commentIntoStruc)
			}
		}

		if tag == "" {
			allPostFinal = append(allPostFinal, postFinalIntoStruc)
		} else {
			if tag == postFinalIntoStruc.Tag {
				allPostFinal = append(allPostFinal, postFinalIntoStruc)
			}

		}

	}
	if len(connectedUser) > 1 {
		xpInt, _ := strconv.Atoi(connectedUser[5])
		homePage = HomePageStruct{
			ConnectedUserXp: xpInt,
			Post:            allPostFinal,
			NbrPost:         len(allPost),
			Comments:        allComment,
			User:            User,
			IsConnected:     true,
		}

	} else {
		homePage = HomePageStruct{
			ConnectedUserXp: 0,
			Post:            allPostFinal,
			NbrPost:         len(allPost),
			Comments:        allComment,
			User:            User,
			IsConnected:     false,
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}
