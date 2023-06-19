package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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