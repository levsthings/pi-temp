# pi-temp

A temperature logger for my Raspberry Pi. It's still under development as I'm exploring native Go GPIO libraries.

### Details

`pi-temp` runs in the background and probes for data every 5 minutes. By default, the data is saved on a daily log file and these daily logs are kept for 7 days 
before deletion. The logging can be swapped to a different system fairly easily like writing to a local or remote database instead of using local log files.

### Python Dependencies

```terminal
sudo apt-get update
sudo apt-get install python3-pip
sudo python3 -m pip install --upgrade pip setuptools wheel
```

### Installation

You can download the latest binary from Github [releases](https://github.com/levsthings/pi-temp/releases). 

### Running

If ran without any flags, the program will write the temperature data once to stdout and exit. If you supply the `--mode log` flag, the program will run
in it's intended background mode writing a log file every 5 minutes infinitely.

See usage:

```terminal
Usage of pi-temp:
  -mode string
    	Expected input: '--mode console' or '--mode log' (default "console")
```

You can run the binary yourself or add it to your startup routine via `systemd` or `rc.local`.The log outputs will automatically go to a directory named 
`.pi-stats`, and will be wherever the binary is run. If you add it to your startup routine, pay attention to the execution context.

### Sample Output

```terminal
18:53:41, Temp: 25.7°C, Humidity: 32.2%
```
