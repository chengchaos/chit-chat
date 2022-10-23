package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chengchaos/chit-chat/utils"
)

type Thread struct {
	Id        int
	UserId    int
	Uuid      string
	Topic     string
	CreatedAt time.Time
}

// CreatedAtDate is method that format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jun 2, 2006 at 3:04pm")
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query(
		"SELECT count(id) FROM posts "+
			"WHERE thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT  id, user_id, thread_id, uuid, body, created_at "+
		"FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			utils.LogError(err, "Close rows get errors!")
		}
	}(rows)

	for rows.Next() {
		post := Post{}
		if err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.ThreadId,
			&post.Uuid,
			&post.Body,
			&post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

// User Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	err := Db.QueryRow("SELECT id, uuid, name, email, created_at "+
		"FROM users WHERE id = $1", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		utils.LogError(err, "What the fuck can I do ?")
	}
	return
}

// User Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	err := Db.QueryRow("SELECT id, uuid, name, email, created_at "+
		"FROM users WHERE id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		utils.LogError(err, "What the fuck can I do?")
	}
	return
}

func CreateThread(user *User, topic string) (conv Thread, err error) {
	return user.CreateThread(topic)
}

func Threads() (threads []Thread, err error) {
	sql := "SELECT a.id, a.user_id, a.uuid, a.topic, a.created_at " +
		"FROM threads a ORDER BY a.created_at DESC "
	rows, err := Db.Query(sql)
	if err != nil {
		utils.LogError(err, "Threads Db Query")
		return
	}
	defer rows.Close()
	for rows.Next() {
		thread := Thread{}
		err = rows.Scan(&thread.Id, &thread.UserId, &thread.Uuid, &thread.Topic, &thread.CreatedAt)

		if err != nil {
			utils.LogError(err, "Threads rows Scan ")
			return
		}
		fmt.Printf("Threads : %v\n", thread)
		threads = append(threads, thread)
	}
	return
}

// ThreadByUUID Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT a.id, a.uuid, a.topic, a.user_id, a.created_at "+
		"FROM threads a WHERE a.uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}
