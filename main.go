package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/speza/runner/pkg/lib"
)

func main() {
	if err := lib.Serve(&Task{}); err != nil {
		log.Fatal(err)
	}
}
