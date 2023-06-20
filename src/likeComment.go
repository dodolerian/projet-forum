package forum

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

// ---- LIKE

var allLikeCommentList []LikeCommentFromDb

func RecuperationLikeComment() {
	allLikeCommentList = nil
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM LIKECOMMENT")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idUser int
		var idComment int
		err = rows.Scan(&idUser, &idComment)

		if err != nil {
			log.Fatal(err)
		}
		likeCommentStruc := LikeCommentFromDb{}
		likeCommentStruc = LikeCommentFromDb{
			IdUser: idUser,
			IdComment: idComment,
		}

		allLikeCommentList = append(allLikeCommentList, likeCommentStruc)

	}
}

func AddLikeComment(db *sql.DB, idUser int, idComment int) {

	records := `INSERT INTO LIKECOMMENT(idUser,idComment) VALUES (?,?)`
	query, err := db.Prepare(records)

	_, err = query.Exec(idUser, idComment)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteLikeComment(db *sql.DB, idUser int, idComment int) {
	records := `DELETE FROM LIKECOMMENT
	WHERE idUser = $1 
	AND idComment = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idUser, idComment)
	if err != nil {
		log.Fatal(err)
	}
}

func LikeOnComment(idUser int, idComment int, allLikeComment []LikeCommentFromDb) bool {
	for i := 0; i < len(allLikeComment); i++ {
		if allLikeComment[i].IdComment == idComment && allLikeComment[i].IdUser == idUser {
			return true
		}
	}
	return false
}

// ------ DISLIKE



var allDislikeCommentList []DislikeCommentFromDb

func RecuperationDislikeComment() {
	allDislikeCommentList = nil
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM DISLIKECOMMENT")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idUser int
		var idComment int
		err = rows.Scan(&idUser, &idComment)

		if err != nil {
			log.Fatal(err)
		}
		dislikeCommentStruc := DislikeCommentFromDb{}
		dislikeCommentStruc = DislikeCommentFromDb{
			IdUser: idUser,
			IdComment: idComment,
		}

		allDislikeCommentList = append(allDislikeCommentList, dislikeCommentStruc)

	}
}

func AddDislikeComment(db *sql.DB, idUser int, idComment int) {

	records := `INSERT INTO DISLIKECOMMENT(idUser,idComment) VALUES (?,?)`
	query, err := db.Prepare(records)

	_, err = query.Exec(idUser, idComment)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDislikeComment(db *sql.DB, idUser int, idComment int) {
	records := `DELETE FROM DISLIKECOMMENT
	WHERE idUser = $1 
	AND idComment = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idUser, idComment)
	if err != nil {
		log.Fatal(err)
	}
}

func DislikeOnComment(idUser int, idComment int, allDislikeComment []DislikeCommentFromDb) bool {
	for i := 0; i < len(allDislikeComment); i++ {
		if allDislikeComment[i].IdComment == idComment && allDislikeComment[i].IdUser == idUser {
			return true
		}
	}
	return false
}


func UseLikeComment(database *sql.DB, likeIdCommentStr string, dislikeIdCommentStr string, connectedUserId int){
	likeIdComment, _ := strconv.Atoi(likeIdCommentStr)
	dislikeIdComment, _ := strconv.Atoi(dislikeIdCommentStr)

	if likeIdCommentStr == "" {
		isLikedComment := LikeOnComment(connectedUserId, dislikeIdComment, allLikeCommentList)
		isDislikedComment := DislikeOnComment(connectedUserId, dislikeIdComment, allDislikeCommentList)

		if isLikedComment {
			DeleteLikeComment(database, connectedUserId, dislikeIdComment)
		}
		if isDislikedComment {
			DeleteDislikeComment(database, connectedUserId, dislikeIdComment)
		} else {
			AddDislikeComment(database, connectedUserId, dislikeIdComment)
		}
	}

	if dislikeIdCommentStr == "" {
		isLikedComment := LikeOnComment(connectedUserId, likeIdComment, allLikeCommentList)
		isDislikedComment := DislikeOnComment(connectedUserId, likeIdComment, allDislikeCommentList)

		if isDislikedComment {
			DeleteDislikeComment(database, connectedUserId, likeIdComment)
		}
		if isLikedComment {
			DeleteLikeComment(database, connectedUserId, likeIdComment)
		} else {
			AddLikeComment(database, connectedUserId, likeIdComment)
		}
	}
}