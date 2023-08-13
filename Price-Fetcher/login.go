package main

import (
	"context"
	"time"
)

type loggingService struct {
	next PriceFetcher
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) {
	defer func(begin time.Time) {

	}(time.Now())
}
