package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// func CreateDatabase() {
//     file, err := os.Create("database.db")
//     if err != nil {
//         log.Fatal(err)
//     }
//     file.Close()

//     database, _ := sql.Open("sqlite3", "database.db")
//     // createTable(database)

//     addUsers(database, "Ankita", "Maudie", "Game Developer", 140000)
//     addUsers(database, "Emiliana", "Alfiya", "Bakend Developer", 120000)
//     addUsers(database, "Emmet", "Brian", "DevOps Developer", 110000)
//     addUsers(database, "Reidun", "Jorge", "Dtabase Developer", 140000)
//     addUsers(database, "Tyrone", "Silvia", "Front-End Developer", 109000)
//     defer database.Close()
//     fetchRecords(database)
// }

// func createTable(db *sql.DB) {
//     users_table := `CREATE TABLE users (
//         id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//         "FirstName" TEXT,
//         "LastName" TEXT,
//         "Dept" TEXT,
//         "Salary" INT);`
//     query, err := db.Prepare(users_table)
//     if err != nil {
//         log.Fatal(err)
//     }
//     query.Exec()
//     fmt.Println("Table created successfully!")
// }

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

func ModifyBDD(db *sql.DB, id int, description string) {
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
