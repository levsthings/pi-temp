package pistats

import (
	"log"
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
// and parses the output from stdin
func GetData() TempData {
	out, err := exec.Command("python3", "./python/temp.py").Output()
	if err != nil {
		log.Fatal("couldn't read from temp.py")
	}

	t := strings.TrimSuffix(string(out), "\n")
	s := strings.Split(t, ",")

	d := [2]float64{}
	for i, str := range s {
		spl := strings.Split(str, "=")
		f, err := strconv.ParseFloat(spl[1], 64)
		if err != nil {
			log.Fatal("couldn't parse output from temp.py")
		}
		d[i] = f
	}

	data := TempData{d[0], d[1]}

	return data
}
