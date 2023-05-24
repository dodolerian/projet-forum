package forum

import (
	"database/sql"
	"log"
	"strconv"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(db *sql.DB , content string, homePage HomePageStruct ){

	records := `INSERT INTO POST(author,like,dislike,date ) VALUES (?,?,?,?,?)`
	query, err := db.Prepare(records)
	idAuthor := homePage.IdAuthor
	idAuthorIntoInt ,_:= strconv.Atoi(idAuthor)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(idAuthorIntoInt, content, 0,0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( content)
}