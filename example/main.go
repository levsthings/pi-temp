package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	consoleMode = "console"
	logMode     = "log"
)

func main() {
	mode := flag.String("mode", consoleMode, "Expected input: '--mode console' or '--mode log'")
	flag.Parse()

	if *mode == logMode {
		for {
			write(format())
			rotate()

			time.Sleep(time.Minute * 5)
		}
	}

	fmt.Print(format())
}

func format() string {
	d := pitemp.GetData()

	var (
		tempFormat = `Temp: %.1f°C, Humidity: %.1f%%`
		logFormat  = "%s, %s\n"
	)

	ti := time.Now().Format("15:04:05")
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
		pitemp.LogFatal(pitemp.ErrorOutput{
			err,
			"couldn't open temp log file",
		})
		os.Exit(1)
	}

	if _, err := f.Write(v); err != nil {
		pitemp.LogFatal(pitemp.ErrorOutput{
			err,
			"couldn't write to temp log file",
		})
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		pitemp.LogFatal(pitemp.ErrorOutput{
			err,
			"couldn't close temp log file",
		})
		os.Exit(1)
	}
}

func rotate() {
	logs, err := ioutil.ReadDir(dir)
	if err != nil {
		pitemp.LogFatal(pitemp.ErrorOutput{
			err,
			"couldn't rotate logs",
		})
		os.Exit(1)
	}

	if len(logs) > maxLogs {
		sort.Slice(logs, func(i, j int) bool {
			t1, _ := time.Parse("02-01-2006", logs[i].Name())
			t2, _ := time.Parse("02-01-2006", logs[j].Name())
			return t1.Before(t2)
		})

		err := os.Remove(logPath + logs[0].Name())
		if err != nil {
			pitemp.LogFatal(pitemp.ErrorOutput{
				err,
				"couldn't delete oldest log file",
			})
			os.Exit(1)
		}
	}
}
