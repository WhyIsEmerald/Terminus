package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/WhyIsEmerald/Terminus/internals/baseconv"
	"github.com/WhyIsEmerald/Terminus/internals/calculator"
	"github.com/WhyIsEmerald/Terminus/internals/units"
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
			{
				Name:    "BaseConvert",
				Aliases: []string{"bc"},
				Action: func(c *cli.Context) error {
					if c.NArg() != 3 {
						return fmt.Errorf("invalid number of arguments, expected <number> <from_base> <to_base>")
					}
					number := c.Args().Get(0)
					fromBaseStr := c.Args().Get(1)
					toBaseStr := c.Args().Get(2)

					fromBase, err := strconv.Atoi(fromBaseStr)
					if err != nil {
						return fmt.Errorf("invalid from_base: %w", err)
					}

					toBase, err := strconv.Atoi(toBaseStr)
					if err != nil {
						return fmt.Errorf("invalid to_base: %w", err)
					}

					result, err := baseconv.BaseConvert(number, fromBase, toBase)
					if err != nil {
						return err
					}

					fmt.Println(result)
					return nil
				},
			},
			{
				Name:    "UnitConv",
				Aliases: []string{"u", "uc"},
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						return fmt.Errorf("unexpected arguments: %v. Did you forget to quote an argument with spaces?", c.Args().Slice())
					}
					measurement := c.String("measurement")
					fromUnit := c.String("from")
					toUnit := c.String("to")
					value := c.Float64("value")

					units.Convert(measurement, fromUnit, toUnit, value)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "measurement",
						Aliases:  []string{"m"},
						Usage:    "The measurement to convert",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "from",
						Aliases:  []string{"f"},
						Usage:    "The unit to convert from",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "to",
						Aliases:  []string{"t"},
						Usage:    "The unit to convert to",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "value",
						Aliases:  []string{"v"},
						Usage:    "The value to convert",
						Required: true,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
