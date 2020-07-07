package main

import (
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func getCommands() []cli.Command {
	// global level flags
	var cpuload float64
	var duration float64
	var cpucore int
	var context *cli.Context
	sampleInterval := 100 * time.Millisecond

	cpuLoadFlags := []cli.Flag{
		cli.Float64Flag{
			Name:        "cpuload",
			Usage:       "Target CPU load 0<cpuload<1",
			Value:       0.1,
			Destination: &cpuload,
		},
		cli.Float64Flag{
			Name:        "duration",
			Usage:       "Duration to run the stress app in Seconds",
			Value:       10,
			Destination: &duration,
		},
		cli.IntFlag{
			Name:        "cpucore",
			Usage:       "Cpu core to stress ",
			Value:       0,
			Destination: &cpucore,
		},
	}
	commands := []cli.Command{
		{
			Name: "cpu",
			Action: func(c *cli.Context) {
				context = c
				runCpuLoader(sampleInterval, cpuload, duration, cpucore)
			},
			Usage:  "load cpu, use --help for more options",
			Flags:  cpuLoadFlags,
			Before: func(_ *cli.Context) error { return nil },
		},
		{
			Name: "memory",
			Action: func(c *cli.Context) {
				context = c
				runMemoryLoader(sampleInterval, cpuload, duration, cpucore)
			},
			Usage:  "load memory, use --help for more options",
			Flags:  cpuLoadFlags,
			Before: func(_ *cli.Context) error { return nil },
		},
	}
	return commands
}

func runCpuLoader(sampleInterval time.Duration, cpuload float64, duration float64, cpu int) {
	controller := NewCpuLoadController(sampleInterval, cpuload)
	monitor := NewCpuLoadMonitor(float64(cpu), sampleInterval)

	actuator := NewCpuLoadGenerator(controller, monitor, time.Duration(duration))
	StartCpuLoadController(controller)
	StartCpuMonitor(monitor)

	RunCpuLoader(actuator)
	StopCpuLoadController(controller)
	StopCpuMonitor(monitor)
}

func runMemoryLoader(sampleInterval time.Duration, cpuload float64, duration float64, cpu int) {
	RunMemoryLoader()
}

func main() {
	app := cli.NewApp()
	app.Name = "Stress"
	app.Usage = "tool to stress test  host !!"
	app.Commands = getCommands()
	app.Version = "0.0.1"
	app.Run(os.Args)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
