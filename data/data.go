package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/chengchaos/chit-chat/utils"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "port=5432 user=chengchao password=Hello-Eko! dbname=chengchao sslmode=disable")
	//还可以是这种方式打开
	//db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// CreateUUID create a random UUID with from RFC 4122
// adapted from http://github.com/un7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return

}

// Encrypt is f function that hash plaintext with SHA-1
func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
}

// GetSession function : <br />
// Checks if the user is logged in has a session, if not err is not nil
func GetSession(w http.ResponseWriter, r *http.Request) (sess utils.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = utils.Session{Uuid: cookie.Value}
		if ok, _ := Check(&sess); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

func Check(session *utils.Session) (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1 ",
		session.Uuid).Scan(&session.Id,
		&session.Uuid,
		&session.Email,
		&session.UserId,
		&session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func DeleteSession(session *utils.Session) (err error) {
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// GetUserBySession method is get a user from session
func GetUserBySession(session *utils.Session) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email create_at FROM users where id = $1 ",
		session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
	return
}
