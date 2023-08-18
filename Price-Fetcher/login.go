package main

import (
	"context"
	"time"
)

type loggingService struct {
	next PriceFetcher
}
