package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/WhyIsEmerald/Terminus/internals/calculator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Terminus",
		Usage: "A collection of useful tools",
		Commands: []*cli.Command{
			{
				Name:    "calculate",
				Aliases: []string{"c"},
				Usage:   "Evaluate a mathematical expression. Remember to quote your expression.",
				Action: func(c *cli.Context) error {
					if c.NArg() > 1 {
						return fmt.Errorf("too many arguments, please quote your expression")
					}
					expression := strings.Join(c.Args().Slice(), " ")
					if expression == "" {
						return cli.ShowSubcommandHelp(c)
					}
					tokens, err := calculator.Tokenize(expression)
					if err != nil {
						return err
					}

					rpnTokens := calculator.ShuntingYard(tokens)
					result, err := calculator.EvaluateRPN(rpnTokens)
					if err != nil {
						return err
					}

					fmt.Println(result)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
