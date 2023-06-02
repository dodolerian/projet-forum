package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type recuperationCommentFromDb struct {
	Id      int
	Author  int
	Content string
	Like    int
	Dislike int
	Date    string
}

var allComment []recuperationPostFromDb

func recuperationComment() []recuperationPostFromDb {
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM COMMENTARY")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var author int
		var content string
		var like int
		var dislike int
		var date string
		err = rows.Scan(&id, &author, &content, &like, &dislike, &date)

		if err != nil {
			log.Fatal(err)
		}
		commentIntoStruc := recuperationCommentFromDb{}
		commentIntoStruc = recuperationCommentFromDb{
			Id:      id,
			Author:  author,
			Content: content,
			Like:    like,
			Dislike: dislike,
			Date:    date,
		}
		allComment = append(allComment, recuperationPostFromDb(commentIntoStruc))

	}
	return allComment
}