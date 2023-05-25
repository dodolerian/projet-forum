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
	   currentTime := time.Now()

	records := `INSERT INTO POST(author,content,like,dislike,date ) VALUES (?,?,?,?,?)`
	query, err := db.Prepare(records)
	idAuthor := homePage.IdAuthor
	idAuthorIntoInt ,_:= strconv.Atoi(idAuthor)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idAuthorIntoInt, content, 0,0, currentTime)
	if err != nil {
		log.Fatal(err)
	}
		fmt.Println("envoie dans la bdd")
}