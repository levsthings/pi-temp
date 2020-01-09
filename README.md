# pi-temp


A temperature and humidity sensor for my Raspberry Pi.

### Details


This project uses DHT22 temperature and humidity sensor to get relative temperature and humidity data. It uses the following wiring:

```asm
VDD <-> 10k OHM <-> DATA
VDD  -> PIN 1
DATA -> PIN 7
NULL
GND  -> PIN 6
```

### Python Dependencies


At the moment, `pi-temp` uses a Python script to interface with the sensor via GPIO as a proof of concept (and the due to accuracy issues with existing Go libraries), so you need to make sure you need to install python dependencies:

```terminal
sudo apt-get install python3-pip
sudo python3 -m pip install --upgrade pip setuptools wheel
```

### Installation


You can download the latest binary from Github [releases](https://github.com/levsthings/pi-temp/releases). 


### Usage


You can import `pi-temp` as a library and use the raw data to feed your own formatter and logger:

```go
package main

import pitemp "github.com/levsthings/pi-temp"

func main() {
    data := pitemp.GetData()
}
```

### Example Program

`pi-temp` comes with an example program that lives in the `example` folder which can be built and used for local logging or getting console outputs.

If `pi-temp` is ran without any flags, the program will write the temperature data once to stdout and exit. If you supply the `--mode log` flag, the program will run
in it's intended background mode writing to a log file every 5 minutes until terminated. By default, the data is saved on a daily log file and these daily logs are kept for 7 days before deletion. The logging can be swapped to a different system fairly easily like writing to a local or remote database instead of using local log files.

See usage:

```terminal
Usage of pi-temp:
  -mode string
    	Expected input: '--mode console' or '--mode log' (default "console")
```

You can run the binary manually or you can add it to your startup routine via `systemd` or `rc.local`.The log outputs will automatically go to a directory named 
`.pi-temp`, and will be wherever the binary is run. If you add it to your startup routine, pay attention to the execution context.

#### Sample Output


```terminal
18:53:41, Temp: 25.7°C, Humidity: 32.2%
```

