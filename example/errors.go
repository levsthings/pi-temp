package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type errorOutput struct {
	Error   error
	Message string
}

func logFatal(e errorOutput) {
	filename := "pitemp-error.log"

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
