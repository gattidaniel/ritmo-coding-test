package ritmo

import (
	"errors"
	"log"
	"time"
)

type Service struct {
	burstLimiter chan time.Time
}

func Start(requestsPerSecond, maxBurst int) Service {
	// We should validate these things, and return an error. I didn't do it to keep the signature define in the main
	if requestsPerSecond < 1 {
		log.Fatalln(errors.New("requestsPerSecond it can't be less than 1"))
	}
	if maxBurst < 0 {
		log.Fatalln(errors.New("maxBurst can't be less than 0"))
	}

	burstLimiter := make(chan time.Time, maxBurst+1)
	burstLimiter <- time.Now()

	go func() {
		for t := range time.Tick(time.Second / time.Duration(requestsPerSecond)) {
			if len(burstLimiter) <= maxBurst {
				burstLimiter <- t
			}
		}
	}()

	return Service{
		burstLimiter: burstLimiter,
	}
}

func (s *Service) Next() {
	<-s.burstLimiter
}
