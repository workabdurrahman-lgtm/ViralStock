package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func runScraper(keyword string) ([]string, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var results []string
	err := chromedp.Run(ctx,
		// a-5-1
	)
	if err != nil {
		log.Println("Scraping failed:", err)
		return nil, err
	}

	log.Println("Scraping successful, results:", results)
	return results, nil
}
