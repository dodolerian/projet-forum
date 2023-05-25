package forum

import (
	"database/sql"
	"log"
	"strconv"
	    "time"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(db *sql.DB , content string, homePage HomePageStruct ){
	    parseTime, err := time.Parse("Jan 02, 2006", "Sep 30, 2021")
    if err != nil {
      panic(err)
    }
   currentTimePArse := parseTime.Format("02, Jan 2006")
	records := `INSERT INTO POST(author,content,like,dislike,date ) VALUES (?,?,?,?,?)`
	query, err := db.Prepare(records)
	idAuthor := homePage.IdAuthor
	idAuthorIntoInt ,_:= strconv.Atoi(idAuthor)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idAuthorIntoInt, content, 0,0, currentTimePArse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("envoie dans la bdd")
	recuperationPost()
}