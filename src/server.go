package forum

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var connectedUser []string

func WebServer() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/create", CreateAccount)
	http.HandleFunc("/", ConnexionAccount)
	http.HandleFunc("/profil", Profil)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 3333 : http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

