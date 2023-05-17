package main

import (
	"database/sql"
	"fmt"
	"forum"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//forum.AddUsers(database, "personne1", "1234", "j'aime les jeux video", "test@gmail.com")
	//forum.ModifyBDD(database, 6, "petit test4")
	WebServer()
}

// func main() {
// 	forum.CreateDatabase()
// 	WebServer()
// }

func WebServer() {

	http.HandleFunc("/", Home)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 8871 : http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type userTable struct {
	Id                int
	Username          string
	Password          string
	ProfilDescription string
	Mail              string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")

	defer database.Close()

	tmpId, tmpUsername, tmpPassword, tmpProfilDescripotion, tmpMail := forum.FetchRecords(database)
	user := userTable{Id: tmpId,
		Username:          tmpUsername,
		Password:          tmpPassword,
		ProfilDescription: tmpProfilDescripotion,
		Mail:              tmpMail}

	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	fmt.Println(r.FormValue("mail"))

	err := tmpl.Execute(w, user)
	if err != nil {
		log.Fatal(err)
	}
}
