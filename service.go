package main

import "context"

type Service interface {
	GetCatFact(context.Context) (*Catfact, error)
}
