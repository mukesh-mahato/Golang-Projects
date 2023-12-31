package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fecth a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements PrceFethcer interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMock = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"MM":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMock[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
