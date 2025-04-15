package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }()
	defer cfg.wg.Done()

	select {
	case <-cfg.ctx.Done():
		return
	default:
		if !cfg.addPageVisit(rawCurrentURL) {
			return
		}

		body, err := getHTML(rawCurrentURL)
		if err != nil {
			log.Println("error reading HTML:", err)
			return
		}
		log.Println("Obtained HTML body from", rawCurrentURL)

		urls, err := getURLsFromHTML(body, cfg.baseURL.Host)
		if err != nil {
			log.Println("error obtaining URLs from HTML:", err)
			return
		}
		log.Println("Found", len(urls), "urls")

		for _, url := range urls {
			if cfg.addPageVisit(url) {
				cfg.wg.Add(1)
				go cfg.crawlPage(url)
			}
		}

		if cfg.isDone {
			cfg.ctxCancel()
			log.Printf("Crawled %d links, stopping", cfg.maxPages)
			return
		}
	}
}

func (cfg *config) addPageVisit(rawURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if checkDomain(rawURL, cfg.baseURL) {
		normalizedURL := normalizeURL(rawURL)

		if _, ok := cfg.pages[normalizedURL]; ok {
			cfg.pages[normalizedURL]++
			return false
		}

		cfg.pages[normalizedURL] = 1
		cfg.isDone = len(cfg.pages) >= cfg.maxPages
		log.Println("stored internal link: ", normalizedURL)

		return true
	}
	log.Println("skipping", rawURL)
	return false
}

func checkDomain(cUrl string, base *url.URL) bool {
	tmp, err := url.Parse(cUrl)
	if err != nil {
		log.Println(err)
	}

	return tmp.Host == base.Host
}
