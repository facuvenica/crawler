package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
	ctx                context.Context
	ctxCancel          context.CancelFunc
	isDone             bool
}

func main() {
	checkArgs(len(os.Args))
	path := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("%s", err)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("%s", err)
	}

	baseURL, err := url.Parse(path)
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg := config{
		baseURL:            baseURL,
		maxPages:           maxPages,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		pages:              make(map[string]int),
		ctx:                ctx,
		ctxCancel:          cancel,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(path)
	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseURL.String())
}

func checkArgs(n int) {
	if n < 4 {
		log.Fatalf("usage: ./crawler URL maxConcurrency maxPages")
	}

	if n > 4 {
		log.Fatalf("too many arguments provided")
	}
}
