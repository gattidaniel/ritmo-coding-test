package main

import (
	"fmt"
	"github.com/mercadolibre/ritmo-coding-test/internal/ritmo"
	"time"
)

func main() {
	fmt.Println("start")
	requestsPerSecond := 10
	maxBurst := 1 // If no jobs are performed during a period of time,
	// those unused jobs can be executed later, with a limit of maxBurst.

	// Method to implement.
	r := ritmo.Start(requestsPerSecond, maxBurst)
	start := time.Now()
	for i := 0; i < 10; i++ {
		if i == 3 {
			time.Sleep(1 * time.Second)
		}
		// Method to implement.
		r.Next()
		job(i, time.Since(start))
	}
	fmt.Println("end")
}

func job(n int, timeDiff time.Duration) {
	fmt.Println("executing job no", n, "after", timeDiff)
}
