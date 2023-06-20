package forum

// Account Structrue
type createAccountStruct struct {
	UsernameError string
	PasswordError string
	MailError     string
}

//Comment Structure
type recuperationCommentFromDb struct {
	IdPost   int
	IdAuthor int
	Content  string
	Like     int
	Dislike  int
	Date     string
}

// Dislike Structure

type DislikeFromDb struct {
	IdUser int
	IdPost int
}

// Home Structure

type HomePageStruct struct {
	ConnectedUserXp int
	Post        []PostStruct
	NbrPost     int
	Comments    []recuperationCommentFromDb
	User        []UserStruct
	IsConnected bool
}

type CommentStruct struct {
	IdPost     int
	IdAuthor   int
	AuthorName string
	Content    string
	Like       int
	Dislike    int
	Date       string
	Image      string
}

type PostStruct struct {
	Id          int
	Author      int
	AuthorName  string
	Content     string
	Like        bool
	Dislike     bool
	Date        string
	Comments    []CommentStruct
	Image       string
	IsImage     bool
	IsConnected bool
	Tag         string
	Xp 			int
}


// Like Structure


type LikeFromDb struct {
	IdUser int
	IdPost int
}

// Post Structure

type recuperationPostFromDb struct {
	Id      int
	Author  int
	Content string
	Like    int
	Dislike int
	Date    string
	Image   []byte
	Tag     string
}


// Profil Structure


type ProfilPageStruct struct {
	Username          string
	ProfilDescription string
	Mail              string
	ConnectedUserXp   int
}

// Recuperation User 

type recuperationUserFromDb struct {
	Id                int
	Username          string
	ProfilDescription string
}

// User struct 
type UserStruct struct {
	Id                int
	Username          string
	ProfilDescription string
	xp				  int
}
