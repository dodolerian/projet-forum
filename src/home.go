package forum

import (
	"database/sql"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"strconv"
	"strings"
)

type HomePageStruct struct {
	Post        []PostStruct
	NbrPost     int
	Comments    []recuperationCommentFromDb
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
	IsConnected bool
	Tag         string
}

func Home(w http.ResponseWriter, r *http.Request) {
	
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	
	homePage := HomePageStruct{}
	// invité := "invité"
	RecuperationLike()
	RecuperationDislike()
	
	allPost = nil
	allPost := recuperationPost()

	if len(connectedUser) == 0 {
		connectedUser = append(connectedUser, "-1")
	}

	connectedUserId, _ := strconv.Atoi(connectedUser[0])
	tag := ""
	/* COMMENTS */
	if r.Method == http.MethodPost {
		censure := r.FormValue("censure")
		if censure !=""{
			fmt.Println("ici")
	}
		IdPost := r.FormValue("idPost")
		tag = r.FormValue("tag")
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

	allComment = nil
	allComment := recuperationComment()

	allCommentOfThisPost := []CommentStruct{}

	allPostFinal := []PostStruct{}
	RecuperationLike()
	RecuperationDislike()

	for i := len(allPost) - 1; i >= 0; i-- {
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

		imgBase64Str := base64.StdEncoding.EncodeToString(allPost[i].Image)

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
			IsConnected: true,
			Tag:         allPost[i].Tag,
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

		if tag == "" {
			allPostFinal = append(allPostFinal, postFinalIntoStruc)
		} else {
			if tag == postFinalIntoStruc.Tag {
				allPostFinal = append(allPostFinal, postFinalIntoStruc)
			}

		}

	}
	if len(connectedUser) > 1 {
		homePage = HomePageStruct{
			Post:        allPostFinal,
			NbrPost:     len(allPost),
			Comments:    allComment,
			IsConnected: true,
		}

	} else {
		homePage = HomePageStruct{
			Post:        allPostFinal,
			NbrPost:     len(allPost),
			Comments:    allComment,
			IsConnected: false,
		}
	}
	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}