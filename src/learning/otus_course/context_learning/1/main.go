package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)
	wg.Add(1)
	go dealLongWithContext(wg, ctx)
	wg.Wait()
}

func dealLongWithContext(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randTime := time.Duration(r.Intn(4000)) * time.Second
	fmt.Printf("Duration %s\n", randTime)

	timer := time.NewTimer(randTime)

	fmt.Println(ctx)
	fmt.Println(time.Now())
	select {
	case <-ctx.Done():
		fmt.Println("Done")
	case <-timer.C:
		fmt.Println("Timer")
	}
}
