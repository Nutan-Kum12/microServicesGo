package main

import (
	"context"
	"fmt"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)
	fact, err := svc.GetCatFact(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println(fact.Fact)
}
