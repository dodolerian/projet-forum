package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddUsers(db *sql.DB, username string, password string, profilDescription string, mail string) {
	records := `INSERT INTO USER(username, password, profilDescription, mail, xp) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(username, password, profilDescription, mail, 0)
	if err != nil {
		log.Fatal(err)
	}
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

func ModifyXpUser(db *sql.DB, id int, xp int) {
	records := `UPDATE USER
	SET  xp = $1
	WHERE id = $2`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(xp, id)
	if err != nil {
		log.Fatal(err)
	}
}

func FetchAllUser(db *sql.DB) ([]int, []string, []string, []string, []string, []int) {
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
	var xp int

	var resId []int
	var resUsername []string
	var resPassword []string
	var resProfilDescription []string
	var resMail []string
	var resXp []int

	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail, &xp)
		resId = append(resId, id)
		resUsername = append(resUsername, username)
		resPassword = append(resPassword, password)
		resProfilDescription = append(resProfilDescription, profilDescription)
		resMail = append(resMail, mail)
		resXp = append(resXp, xp)
	}
	return resId, resUsername, resPassword, resProfilDescription, resMail, resXp
}

func FetchUserWithName(db *sql.DB, name string) (int, string, string, string, string, int) {
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
	var xp int

	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail, &xp)
	}
	return id, username, password, profilDescription, mail, xp
}

func FetchUserWithMail(db *sql.DB, mailEnter string) (int, string, string, string, string, int) {
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
	var xp int

	for record.Next() {
		record.Scan(&id, &username, &password, &profilDescription, &mail, &xp)
	}
	return id, username, password, profilDescription, mail, xp
}

func FetchUserWithId(db *sql.DB, id string) (int, string, string, string, string, int) {
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
	var xp int

	for record.Next() {
		record.Scan(&idFetch, &username, &password, &profilDescription, &mail, &xp)
	}
	return idFetch, username, password, profilDescription, mail, xp
}