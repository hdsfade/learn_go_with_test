package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int,err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	sleeper := &ConfigurableSleeper{1*time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}