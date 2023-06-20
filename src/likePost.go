package forum

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

// ---- LIKE

var allLikeList []LikeFromDb

func RecuperationLike() {
	allLikeList = nil
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM LIKE")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idUser int
		var idPost int
		err = rows.Scan(&idUser, &idPost)

		if err != nil {
			log.Fatal(err)
		}
		likeStruc := LikeFromDb{}
		likeStruc = LikeFromDb{
			IdUser: idUser,
			IdPost: idPost,
		}

		allLikeList = append(allLikeList, likeStruc)

	}
}

func AddLike(db *sql.DB, idUser int, idPost int) {

	records := `INSERT INTO LIKE(idUser,idPost) VALUES (?,?)`
	query, err := db.Prepare(records)

	_, err = query.Exec(idUser, idPost)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteLike(db *sql.DB, idUser int, idPost int) {
	records := `DELETE FROM LIKE
	WHERE idUser = $1 
	AND idPost = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idUser, idPost)
	if err != nil {
		log.Fatal(err)
	}
}

func LikeOnPost(idUser int, idPost int, allLike []LikeFromDb) bool {
	for i := 0; i < len(allLike); i++ {
		if allLike[i].IdPost == idPost && allLike[i].IdUser == idUser {
			return true
		}
	}
	return false
}

// ------ DISLIKE



var allDislikeList []DislikeFromDb

func RecuperationDislike() {
	allDislikeList = nil
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM DISLIKE")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idUser int
		var idPost int
		err = rows.Scan(&idUser, &idPost)

		if err != nil {
			log.Fatal(err)
		}
		dislikeStruc := DislikeFromDb{}
		dislikeStruc = DislikeFromDb{
			IdUser: idUser,
			IdPost: idPost,
		}

		allDislikeList = append(allDislikeList, dislikeStruc)

	}
}

func AddDislike(db *sql.DB, idUser int, idPost int) {

	records := `INSERT INTO DISLIKE(idUser,idPost) VALUES (?,?)`
	query, err := db.Prepare(records)

	_, err = query.Exec(idUser, idPost)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDislike(db *sql.DB, idUser int, idPost int) {
	records := `DELETE FROM DISLIKE
	WHERE idUser = $1 
	AND idPost = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idUser, idPost)
	if err != nil {
		log.Fatal(err)
	}
}

func DislikeOnPost(idUser int, idPost int, allDislike []DislikeFromDb) bool {
	for i := 0; i < len(allDislike); i++ {
		if allDislike[i].IdPost == idPost && allDislike[i].IdUser == idUser {
			return true
		}
	}
	return false
}


func UseLike(database *sql.DB, likeIdPostStr string, dislikeIdPostStr string, connectedUserId int){
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