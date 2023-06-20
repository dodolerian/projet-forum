package forum

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
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
		var date string
		var idComment int
		err = rows.Scan(&idPost, &idAuthor, &content, &date, &idComment)

		if err != nil {
			log.Fatal(err)
		}
		commentIntoStruc := recuperationCommentFromDb{}
		commentIntoStruc = recuperationCommentFromDb{
			IdPost:   idPost,
			IdAuthor: idAuthor,
			Content:  content,
			Date:     date,
			IdComment: idComment,
		}
		allComment = append(allComment, commentIntoStruc)

	}
	return allComment
}

func AddComment(db *sql.DB, content string, idAuthor string, idPost string) {
	parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")
	if err != nil {
		panic(err)
	}
	currentTimePArse := parseTime.Format("02, Jan 2006")
	records := `INSERT INTO COMMENTARY(idPost,idAuthor,content,date ) VALUES (?,?,?,?)`
	query, err := db.Prepare(records)
	idPostStr := idPost
	idPostIntoInt, _ := strconv.Atoi(idPostStr)
	idAuthorStr := idAuthor
	idAuthorIntoInt, _ := strconv.Atoi(idAuthorStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idPostIntoInt, idAuthorIntoInt, content, currentTimePArse)
	if err != nil {
		log.Fatal(err)
	}
}