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

### Installation & Running

**TODO: Add installation methods**


### Sample Output

```terminal
18:53:41, Temp: 25.7°C, Humidity: 32.2%
```
