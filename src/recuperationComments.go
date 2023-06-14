package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type recuperationCommentFromDb struct {
	IdPost   int
	IdAuthor int
	Content  string
	Like     int
	Dislike  int
	Date     string
}

var allComment []recuperationCommentFromDb

func recuperationComment() []recuperationCommentFromDb {
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM COMMENTARY")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idPost int
		var idAuthor int
		var content string
		var like int
		var dislike int
		var date string
		err = rows.Scan(&idPost, &idAuthor, &content, &like, &dislike, &date)

		if err != nil {
			log.Fatal(err)
		}
		commentIntoStruc := recuperationCommentFromDb{}
		commentIntoStruc = recuperationCommentFromDb{
			IdPost:   idPost,
			IdAuthor: idAuthor,
			Content:  content,
			Like:     like,
			Dislike:  dislike,
			Date:     date,
		}
		allComment = append(allComment, commentIntoStruc)

	}
	return allComment
}
