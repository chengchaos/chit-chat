package data

import (
	"log"
	"testing"
)

func TestThreads(t *testing.T) {
	threads, err := Threads()

	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		for _, thr := range threads {
			log.Printf("threads : %v\n", thr)
		}
	}
}
