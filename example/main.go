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
	maxLogs = 7

	consoleMode = "console"
	logMode     = "log"
)

func main() {
	mode := flag.String("mode", consoleMode, "Expected input: '--mode console' or '--mode log'")
	flag.Parse()

	if *mode == logMode {
		for {
			d, err := pitemp.GetData()
			if err != nil {
				logError(errorOutput{
					err,
					"couldn't get temp data from pi-temp",
				})
			}
			write(format(d))
			rotate()

			time.Sleep(time.Minute * 5)
		}
	}

	d, err := pitemp.GetData()
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't get temp data from pi-temp",
		})
	}
	fmt.Print(format(d))
}

func format(d *pitemp.TempData) string {
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
		err := os.Mkdir(dir, 0755)
		if err != nil {
			logError(errorOutput{
				Error:   err,
				Message: "couldn't create log folder",
			})
		}
	}

	f, err := os.OpenFile(logPath+t, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't open temp log file",
		})
	}

	if _, err := f.Write(v); err != nil {
		logError(errorOutput{
			err,
			"couldn't write to temp log file",
		})
	}

	if err := f.Close(); err != nil {
		logError(errorOutput{
			err,
			"couldn't close temp log file",
		})
	}
}

func rotate() {
	logs, err := ioutil.ReadDir(dir)
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't rotate logs",
		})
	}

	if len(logs) > maxLogs {
		sort.Slice(logs, func(i, j int) bool {
			t1, _ := time.Parse("02-01-2006", logs[i].Name())
			t2, _ := time.Parse("02-01-2006", logs[j].Name())
			return t1.Before(t2)
		})

		err := os.Remove(logPath + logs[0].Name())
		if err != nil {
			logError(errorOutput{
				err,
				"couldn't delete oldest log file",
			})
		}
	}
}
