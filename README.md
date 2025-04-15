## Crawler

Very basic web crawler written in Golang as part of a bootdev course. You can clone the repository locally and run it from the command line

> go run .

Or it can be installed with the command install
> go install github.com/facuvenica/crawler

Usage is as simple as calling it with three params
> crawler [URL] [maxConcurrency] [maxPages]

* URL is the website to crawl
* maxConcurrency is a limit on the number of goroutines running simultaneously
* maxPages sets a limit for the amount of internal pages that will be crawled

At the end it will return a neat report of the amount of references to each internal link