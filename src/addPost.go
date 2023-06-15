package forum

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(db *sql.DB, content string, idAutor string, imageFile []byte, tag string) {

	parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")
	if err != nil {
		panic(err)
	}
	currentTimePArse := parseTime.Format("02, Jan 2006")
	records := `INSERT INTO POST(author,content,like,dislike,date,image,tag ) VALUES (?,?,?,?,?,?,?)`
	query, err := db.Prepare(records)
	idAuthor := idAutor
	idAuthorIntoInt, _ := strconv.Atoi(idAuthor)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idAuthorIntoInt, content, 0, 0, currentTimePArse, imageFile, tag)
	if err != nil {
		log.Fatal(err)
	}
}
