package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddUsers(db *sql.DB, username string, password string, profilDescription string, mail string) {
	records := `INSERT INTO USER(username, password, profilDescription, mail) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(username, password, profilDescription, mail)
	if err != nil {
		log.Fatal(err)
	}
}

func FetchAllUser(db *sql.DB) ([]int, []string, []string, []string, []string) {
	record, err := db.Query("SELECT * FROM USER")
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	var id int
	var username string
	var password string
	var profilDescription string
	var mail string

	var resId []int
	var resUsername []string
	var resPassword []string
	var resProfilDescription []string
	var resMail []string
	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail)
		resId = append(resId, id)
		resUsername = append(resUsername, username)
		resPassword = append(resPassword, password)
		resProfilDescription = append(resProfilDescription, profilDescription)
		resMail = append(resMail, mail)
	}
	return resId, resUsername, resPassword, resProfilDescription, resMail
}

func ModifyDescriptionUser(db *sql.DB, id int, description string) {
	records := `UPDATE USER
	SET  profilDescription = $1
	WHERE id = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(description, id)
	if err != nil {
		log.Fatal(err)
	}
}

func FetchUserWithName(db *sql.DB, name string) (int, string, string, string, string) {
	record, err := db.Query("SELECT * FROM USER WHERE username = '" + name + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	var id int
	var username string
	var password string
	var profilDescription string
	var mail string

	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail)
	}
	return id, username, password, profilDescription, mail
}

func FetchUserWithMail(db *sql.DB, mailEnter string) (int, string, string, string, string) {
	record, err := db.Query("SELECT * FROM USER WHERE mail = '" + mailEnter + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	var id int
	var username string
	var password string
	var profilDescription string
	var mail string

	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail)
	}
	return id, username, password, profilDescription, mail
}

func FetchUserWithId(db *sql.DB, id string) (int, string, string, string, string) {
	record, err := db.Query("SELECT * FROM USER WHERE id = '" + id + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	var idFetch int
	var username string
	var password string
	var profilDescription string
	var mail string

	for record.Next() {
		record.Scan(&idFetch, &username, &password, &profilDescription, &mail)
	}
	return idFetch, username, password, profilDescription, mail
}
