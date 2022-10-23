package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chengchaos/chit-chat/utils"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// CreateSession create a new session for an existing user
func (user *User) CreateSession() (session utils.Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) " +
		" values ($1, $2, $3, $4) " +
		" returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows := stmt.QueryRow(CreateUUID(), user.Email, user.Id, user.CreatedAt)
	err = rows.Scan(&session.Id,
		&session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		fmt.Printf("Create Session %s \n", err.Error())
	}

	return
}

func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) " +
		"returning id, uuid, created_at "
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			utils.LogError(err)
		}
	}(stmt)
	// use QueryRow to return a row and scan the returned id into the User struct
	row := stmt.QueryRow(CreateUUID(), user.Name, user.Email,
		Encrypt(user.Password), time.Now())
	fmt.Printf("row => %v\n", *row)
	err = row.Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// CreateThread is a method that to create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic, user_id, created_at) " +
		" values ($1, $2, $3, $4) " +
		" returning id, uuid, topic, user_id, created_at "
	stat, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stat.Close()

	err = stat.QueryRow(CreateUUID(), topic, user.Id, time.Now()).
		Scan(&conv.Id,
			&conv.Uuid,
			&conv.Topic,
			&conv.UserId,
			&conv.CreatedAt)
	return
}

func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.Logger.Println("Close statement ", err)
		}
	}(stmt)
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(CreateUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at "+
		"FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
