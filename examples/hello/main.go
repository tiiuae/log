package main

import (
	"context"

	"github.com/tiiuae/log"
)

func main() {
	ctx := context.Background()
	log.Notice(ctx, "Start up")

	log.Debug(ctx, "Logging with attributes", log.A("some_integer", 5), log.A("some_boolean", false))

	log.Notice(ctx, "Shut down")
}
