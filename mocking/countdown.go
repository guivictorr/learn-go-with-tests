package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{time.Sleep, 1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
