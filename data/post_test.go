package data

import (
	"log"
	"testing"
)

func TestPostSave(t *testing.T) {

	post := &Post{
		UserId:   1,
		ThreadId: 1,
		Body:     "Hello world!",
	}

	err := post.Save()

	if err != nil {
		log.Fatalf("err : %v\n", err.Error())
	} else {
		log.Printf("post: %v\n", *post)
	}
}
