# **Stress test CPU, Memory, Disk and IO with GO** #

* Stress testing of CPU, RAM, Disk and IO with go.
* CPULoad test is a port of [CPULoadGenerator](https://github.com/GaetanoCarlucci/CPULoadGenerator)
* Forked from [dhoomakethu/stress](https://github.com/dhoomakethu/stress)

## Build ##

```bash
git clone https://github.com/yene/stress.git
cd stress/
go build
```

### Compile for Linux ###
```bash
GOOS=linux GOARCH=amd64 go build
```

## Usage ##
### General usage ###
./stress <command> <options>
```bash

./stress --help
NAME:
   Stress - tool to stress test  host !!

USAGE:
   stress [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   cpu		load cpu , use --help for more options
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
### To load CPU to a particular value ###
```bash
./stress cpu --help
NAME:
   stress cpu - load cpu , use --help for more options

USAGE:
   stress cpu [command options] [arguments...]

OPTIONS:
   --cpuload "0.1"	Target CPU load 0<cpuload<1
   --duration "10"	Duration to run the stress app in Seconds
   --cpucore "0"	Cpu core to stress
```
### Examples ###
To load CPU core 1 to 50% for a duration of 10 seconds

```bash
./stress cpu --cpuload 0.5 --duration 10 --cpu 0
```
