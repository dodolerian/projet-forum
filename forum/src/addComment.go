package forum

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddComment(db *sql.DB, content string, idAuthor string, idPost string) {
	parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")
	if err != nil {
		panic(err)
	}
	currentTimePArse := parseTime.Format("02, Jan 2006")
	records := `INSERT INTO COMMENTARY(idPost,idAuthor,content,like,dislike,date ) VALUES (?,?,?,?,?,?)`
	query, err := db.Prepare(records)
	idPostStr := idPost
	idPostIntoInt, _ := strconv.Atoi(idPostStr)
	idAuthorStr := idAuthor
	idAuthorIntoInt, _ := strconv.Atoi(idAuthorStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idPostIntoInt, idAuthorIntoInt, content, 0, 0, currentTimePArse)
	if err != nil {
		log.Fatal(err)
	}
}
