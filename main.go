package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/nicholasjackson/periph-gpio-simulator/host/rpi"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lmicroseconds)
	logger.Println("Hello World")

	// Load all drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	p14 := PinCycle{
		Pin:     rpi.SO_51,
		Running: false,
	}

	p15 := PinCycle{
		Pin:     rpi.SO_53,
		Running: false,
	}

	p18 := PinCycle{
		Pin:     rpi.SO_63,
		Running: false,
	}

	p23 := PinCycle{
		Pin:     rpi.SO_77,
		Running: false,
	}

	p24 := PinCycle{
		Pin:     rpi.SO_81,
		Running: false,
	}

	p25 := PinCycle{
		Pin:     rpi.SO_83,
		Running: false,
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("mode") == "on" {
			p14.Cycle()
			p15.Cycle()
			p18.Cycle()
			p23.Cycle()
			p24.Cycle()
			p25.Cycle()
		} else {
			p14.Stop()
			p15.Stop()
			p18.Stop()
			p23.Stop()
			p24.Stop()
			p25.Stop()
		}
	})

	http.ListenAndServe(":9000", nil)

	for {
	}
}

type PinCycle struct {
	Pin     gpio.PinIO
	Running bool
}

func (f *PinCycle) Cycle() {
	go func() {
		f.Running = true
		state := gpio.High

		for f.Running {
			f.Pin.Out(state)

			sleepDuration := rand.Intn(1000-300) + 300
			time.Sleep(time.Duration(sleepDuration) * time.Millisecond)

			// flip the state
			if state == gpio.High {
				state = gpio.Low
			} else {
				state = gpio.High
			}
		}
	}()
}

func (f *PinCycle) Stop() {
	f.Running = false
}
