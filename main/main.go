package main

import forum "forum/src"

func main() {
	//forum.AddUsers(database, "personne1", "1234", "j'aime les jeux video", "test@gmail.com")
	//forum.ModifyBDD(database, 6, "petit test4")
<<<<<<< HEAD
	forum.WebServer()
=======
	WebServer()
}

func WebServer() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/create", CreateAccount)
	http.HandleFunc("/", ConnexionAccount)
	http.HandleFunc("/profil", Profil)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 8871 : http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type createAccountStruct struct {
	UsernameError string
	PasswordError string
	MailError     string
}

type HomePageStruct struct {
	Username          string
	ProfilDescription string
	Mail              string
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/createAccount.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()

	_, tmpUsername, _, _, tmpMail := forum.FetchAllUser(database)
	accountPage := createAccountStruct{}

	usernameForm := r.FormValue("username")
	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")

	if usernameForm != "" && passwordForm != "" && mailForm != "" {

		if ContainsStringArray(tmpUsername, usernameForm) {
			accountPage = createAccountStruct{UsernameError: "nom déja utilisé"}

		} else if len(passwordForm) < 5 {
			accountPage = createAccountStruct{PasswordError: "mot de passe trop court"}

		} else if ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail déja utilisé"}

		} else {
			//hash
			forum.AddUsers(database, usernameForm, passwordForm, "", mailForm)
			id, username, password, profilDescription, mail := forum.FetchUserWithName(database, usernameForm)
			connectedUser = append(connectedUser, strconv.Itoa(id), username, password, profilDescription, mail)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnexionAccount(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/connexionAccount.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()

	_, _, _, _, tmpMail := forum.FetchAllUser(database)
	accountPage := createAccountStruct{}

	passwordForm := r.FormValue("password")
	mailForm := r.FormValue("mail")

	if passwordForm != "" && mailForm != "" {
		if !ContainsStringArray(tmpMail, mailForm) {
			accountPage = createAccountStruct{MailError: "adresse mail pas trouvé"}
		} else {
			id, username, password, profilDescription, mail := forum.FetchUserWithMail(database, mailForm)
			//dehash
			if passwordForm != password {
				accountPage = createAccountStruct{PasswordError: "mot de passe faux"}
			} else {
				connectedUser = append(connectedUser, strconv.Itoa(id), username, password, profilDescription, mail)
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}
	}

	err := tmpl.Execute(w, accountPage)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))

	homePage := HomePageStruct{}
	if len(connectedUser) > 0 {
		homePage = HomePageStruct{
			Username:          connectedUser[1],
			ProfilDescription: connectedUser[3],
			Mail:              connectedUser[4],
		}
	}

	err := tmpl.Execute(w, homePage)
	if err != nil {
		log.Fatal(err)
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/profil.html"))

	database, _ := sql.Open("sqlite3", "./database/forumBDD.db")
	defer database.Close()

	description := r.FormValue("description")

	// si le form est rempli alors change la valeur, empache qu'elle soit vide
	if description != "" {
		id, _ := strconv.Atoi(connectedUser[0])
		forum.ModifyDescriptionUser(database, id, description)
	}

	// reprend l'utilisateur modifié a chaque fois
	idRefresh, username, password, profilDescription, mail := forum.FetchUserWithId(database, connectedUser[0])
	connectedUser = nil
	connectedUser = append(connectedUser, strconv.Itoa(idRefresh), username, password, profilDescription, mail)

	profilPage := HomePageStruct{
		Username:          connectedUser[1],
		ProfilDescription: connectedUser[3],
		Mail:              connectedUser[4],
	}

	err := tmpl.Execute(w, profilPage)
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
>>>>>>> dev-back
}
