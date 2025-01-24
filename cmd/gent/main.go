package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gent",
		Usage: "Generate Go types for Tree-sitter nodes",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("Hello, world!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
