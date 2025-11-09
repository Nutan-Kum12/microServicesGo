package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	Next Service
}

func NewLoggingService(next Service) Service {
	if next == nil {
		panic("next service cannot be nil")
	}
	return &LoggingService{Next: next}
}
func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("GetCatFact took %v\n", time.Since(start))
	}(time.Now())
	return s.Next.GetCatFact(ctx)
}
