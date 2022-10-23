package data

import "time"

type Post struct {
	Id        int
	UserId    int
	ThreadId  int
	Uuid      string
	Body      string
	CreatedAt time.Time
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jun 2, 2006 at 3:04pm")
}

func (post *Post) Save() (err error) {
	statement := "INSERT INTO posts (uuid, body, user_id, thread_id , created_at) values ($1, $2, $3, $4, $5) " +
		"returning id, created_at "
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(CreateUUID(),
		post.Body,
		post.UserId,
		post.ThreadId,
		time.Now()).Scan(&post.Id, &post.CreatedAt)
	return
}
