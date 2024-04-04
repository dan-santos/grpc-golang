package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(nxt PriceFetcher) PriceFetcher {
	return &metricService{
		next: nxt,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus")
	// TODO: push metrics here, opentelemetry, prometheus, etc
	return s.next.FetchPrice(ctx, ticker)
}