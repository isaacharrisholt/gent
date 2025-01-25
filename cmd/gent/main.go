package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/isaacharrisholt/gent"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gent",
		Usage: "Go ENhancements for Tree-sitter",
		Commands: []*cli.Command{
			generateCommand(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateCommand() *cli.Command {
	return &cli.Command{
		Name:                   "generate",
		Usage:                  "Generate Go types from Tree-sitter node-types.json files",
		UsageText:              "gent generate [OPTIONS] <PATH TO NODE-TYPES.JSON>",
		Aliases:                []string{"gen"},
		Action:                 generateCommandAction,
		EnableShellCompletion:  true,
		Suggest:                true,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "package",
				Aliases: []string{"p"},
				Usage:   "Specify the `PACKAGE` name used for the generated code.",
				Value:   "node_types",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Specify the `OUTPUT` file path used for the generated code. If not specified, the output will be written to stdout.",
			},
		},
	}
}

func generateCommandAction(ctx context.Context, cmd *cli.Command) error {
	if len(cmd.Args().Slice()) == 0 {
		return cli.ShowSubcommandHelp(cmd)
	}

	filePath := cmd.Args().First()
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failed to read from %s: %w", filePath, err)
	}

	generator := gent.NewGenerator(gent.GeneratorOptions{
		PackageName: cmd.String("package"),
	})
	output, err := generator.Generate(fileContent)
	if err != nil {
		return fmt.Errorf("Failed to generate Go code: %w", err)
	}

	if cmd.String("output") != "" {
		if err := os.WriteFile(cmd.String("output"), []byte(output), 0644); err != nil {
			return fmt.Errorf("Failed to write to %s: %w", cmd.String("output"), err)
		}
		return nil
	}

	fmt.Println(output)

	return nil
}
