package main

import (
	_ "coelhino/cli"
	"coelhino/context"

	"log"
)

func main() {
	// cli.Execute()
	ctx := context.NewProjectCtx()

	log.Printf("%v\n", ctx.PublicStr)
}
