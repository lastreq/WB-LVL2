package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_FewChan(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	var minTime float64 = 3
	<-fewChan(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	durSec := time.Since(start).Seconds()
	if durSec < (minTime + 0.2) {
		fmt.Println("OK")
	}

}
