package forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type recuperationUserFromDb struct {
	Id                int
	Username          string
	ProfilDescription string
}

var allUsers []recuperationUserFromDb

func recuperationUser() []recuperationUserFromDb {
	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()
	rows, err := database.Query("SELECT * FROM USER")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var password string
		var profilDescription string
		var mail string
		var xp string
		err = rows.Scan(&id, &username, &password, &profilDescription, &mail, &xp)

		if err != nil {
			log.Fatal(err)
		}
		userIntoStruc := recuperationUserFromDb{}
		userIntoStruc = recuperationUserFromDb{
			Id:                id,
			Username:          username,
			ProfilDescription: profilDescription,
		}
		allUsers = append(allUsers, userIntoStruc)

	}
	return allUsers
}
