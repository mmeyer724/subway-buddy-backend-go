package main

import (
	"buddy/feed"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	for ; true; <-ticker.C {
		feed.PullFeed()
	}
}
