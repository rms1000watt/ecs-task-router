package router

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Router is the router logic
func Router(cfg Config) {
	fmt.Println("Start: Router")
	defer fmt.Println("Done: Router")

	fmt.Println()

	log.Debug(cfg)
}
