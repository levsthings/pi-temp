package pitemp

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	perms    = 0770
	filename = "pitemp.log"
)

type ErrorOutput struct {
	Error   error
	Message string
}

// LogFatal is a naive function that logs crashes
func LogFatal(e ErrorOutput) {

	t := time.Now().Format("15:04:05")
	s := fmt.Sprintf("time=%q, level=fatal, error=%q, msg=%q\n", t, e.Error, e.Message)
	b := []byte(s)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perms)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
