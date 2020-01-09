package pitemp

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// TempData contains temperature and humidity values
type TempData struct {
	Temperature float64
	Humidity    float64
}

// GetData invokes a python script that talks to a GPIO line
// and parses the output from stdout
func GetData() TempData {
	out, err := exec.Command("python3", "-c", py).Output()
	if err != nil {
		LogFatal(ErrorOutput{
			err,
			"couldn't read from python script",
		})
		os.Exit(1)
	}

	t := strings.TrimSuffix(string(out), "\n")
	s := strings.Split(t, ",")

	d := [2]float64{}
	for i, str := range s {
		spl := strings.Split(str, "=")
		f, err := strconv.ParseFloat(spl[1], 64)
		if err != nil {
			LogFatal(ErrorOutput{
				err,
				"couldn't parse output from python script",
			})

			os.Exit(1)
		}
		d[i] = f
	}

	data := TempData{d[0], d[1]}

	return data
}
