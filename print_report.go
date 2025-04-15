package main

import (
	"fmt"
	"slices"
	"strings"
)

type data struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
  REPORT for %s
=============================
`, baseURL)

	for _, link := range sortPages(pages) {
		fmt.Printf("Found %d internal links to %s\n", link.count, link.url)
	}
}

func sortPages(toSort map[string]int) (sorted []data) {
	allData := make([]data, 0, len(toSort))

	for url, count := range toSort {
		allData = append(allData, data{url: url, count: count})
	}

	// Sort by count in descending order, then alphabetically
	slices.SortFunc(allData, func(a, b data) int {
		if a.count != b.count {
			// The value is negative when a is higher
			return b.count - a.count
		}
		// Now we compare alphabetical order
		return strings.Compare(a.url, b.url)
	})

	return allData
}
