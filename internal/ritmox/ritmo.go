package ritmox

// This package is much more readable, and easy to maintain. Although it didn't cover exactly what have been asked
// because it starts with burst. I leave it here in case you decide it is better.

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"
)

type Service struct {
	limiter *rate.Limiter
}

func Start(requestsPerSecond, maxBurst int) Service {
	// We should validate these things, and return an error. I didn't do it to keep the signature define in the main
	if requestsPerSecond < 1 {
		log.Fatalln(errors.New("requestsPerSecond it can't be less than 1"))
	}
	if maxBurst < 1 {
		log.Fatalln(errors.New("maxBurst can't be less than 1"))
	}
	return Service{
		limiter: rate.NewLimiter(rate.Every(time.Second/time.Duration(requestsPerSecond)), maxBurst),
	}
}

func (s *Service) Next() {
	// It would be nice to receive context by parameter, I didn't do it to keep the signature define in the main
	if err := s.limiter.Wait(context.Background()); err != nil {
		log.Fatalln(fmt.Errorf("fail while waiting: %v", err))
	}
}
