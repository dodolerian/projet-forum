package forum

import (
	"database/sql"
	"log"
	"strconv"
	"time"
"fmt"
	_ "github.com/mattn/go-sqlite3"
)
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
		var image []byte
		var tag string
		err = rows.Scan(&id, &author, &content, &like, &dislike, &date, &image, &tag)

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
			Image:   image,
			Tag:     tag,
		}

		allPost = append(allPost, postIntoStruc)

	}
	return allPost
}
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