package forum

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func CreateDatabase() {
    fmt.Println("icccccccccccccccccc")
    file, err := os.Create("database.db")
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    database, _ := sql.Open("sqlite3", "database.db")
    // createTable(database)

    addUsers(database, "Ankita", "Maudie", "Game Developer", 140000)
    addUsers(database, "Emiliana", "Alfiya", "Bakend Developer", 120000)
    addUsers(database, "Emmet", "Brian", "DevOps Developer", 110000)
    addUsers(database, "Reidun", "Jorge", "Dtabase Developer", 140000)
    addUsers(database, "Tyrone", "Silvia", "Front-End Developer", 109000)
    defer database.Close()
    fetchRecords(database)
}

func createTable(db *sql.DB) {
    users_table := `CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "FirstName" TEXT,
        "LastName" TEXT,
        "Dept" TEXT,
        "Salary" INT);`
    query, err := db.Prepare(users_table)
    if err != nil {
        log.Fatal(err)
    }
    query.Exec()
    fmt.Println("Table created successfully!")
}

func addUsers(db *sql.DB, FirstName string, LastName string, Dept string, Salary int) {
    records := `INSERT INTO users(FirstName, LastName, Dept, Salary) VALUES (?, ?, ?, ?)`
    query, err := db.Prepare(records)
    if err != nil {
        log.Fatal(err)
    }
    _, err = query.Exec(FirstName, LastName, Dept, Salary)
    if err != nil {
        log.Fatal(err)
    }
}

func fetchRecords(db *sql.DB) {
    record, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer record.Close()
    for record.Next() {
        var id int
        var FirstName string
        var LastName string
        var Dept string
        var Salary int
        record.Scan(&id, &FirstName, &LastName, &Dept, &Salary)
        fmt.Printf("User: %d %s %s %s %d", id, FirstName, LastName, Dept, Salary)
        fmt.Println()
    }
}