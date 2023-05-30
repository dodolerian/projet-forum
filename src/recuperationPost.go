package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type recuperationPostFromDb struct {
	Id      int
	Author  int
	Content string
	Like    int
	Dislike int
	Date    string
}

var allPost []recuperationPostFromDb

func recuperationPost() []recuperationPostFromDb {
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM POST")
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
		postIntoStruc := recuperationPostFromDb{}
		postIntoStruc = recuperationPostFromDb{
			Id:      id,
			Author:  author,
			Content: content,
			Like:    like,
			Dislike: dislike,
			Date:    date,
		}
		allPost = append(allPost, postIntoStruc)

	}
	return allPost
}
