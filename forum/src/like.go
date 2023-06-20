<<<<<<< HEAD:forum/src/like.go
package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type LikeFromDb struct {
	IdUser int
	IdPost int
}

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
=======
package forum

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


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
>>>>>>> 3b071879a73e125a9cbbab3d8374d2d25a1e592f:src/like.go
