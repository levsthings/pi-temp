package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	pitemp "github.com/levsthings/pi-temp"
)

const (
	dir     = ".pi-temp"
	logPath = dir + "/"
	perms   = 0770
	maxLogs = 7
)

func main() {
	for {
		f := format()

		write(f)
		rotate()

		time.Sleep(time.Minute * 5)
	}
}

func format() string {
	d := pitemp.GetData()

	var (
		tempFormat = `Temp: %.1f°C Humidity: %.1f%%`
		timeFormat = `%d:%d:%d`
		logFormat  = "%s, %s\n"
	)

	h, m, s := time.Now().Clock()

	ti := fmt.Sprintf(timeFormat, h, m, s)
	te := fmt.Sprintf(tempFormat, d.Temperature, d.Humidity)
	log := fmt.Sprintf(logFormat, ti, te)

	return log
}

func write(d string) {
	t := time.Now().Format("02-01-2006")
	v := []byte(d)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, perms)
	}

	f, err := os.OpenFile(logPath+t, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perms)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(v); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func rotate() {
	logs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("couldn't rotate logs", err)
	}

	if len(logs) > maxLogs {
		sort.Slice(logs, func(i, j int) bool {
			t1, _ := time.Parse("02-01-2006", logs[i].Name())
			t2, _ := time.Parse("02-01-2006", logs[j].Name())
			return t1.Before(t2)
		})

		err := os.Remove(logPath + logs[0].Name())
		if err != nil {
			log.Println("error deleting oldest log", err)
		}
	}
}
