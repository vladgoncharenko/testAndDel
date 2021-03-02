package main

import "time"

func main() {
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-timer.C:
			println("every 2 second")

		}
	}
}
