<<<<<<< HEAD:forum/src/server.go
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

func ContainsStringArray(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}
	return false
}
=======
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

>>>>>>> 3b071879a73e125a9cbbab3d8374d2d25a1e592f:src/server.go
