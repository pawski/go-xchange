package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli/v2"
	_ "go.uber.org/automaxprocs"

	"github.com/pawski/go-xchange/internal/application/command"
	"github.com/pawski/go-xchange/internal/logger"
)

func main() {
	app := cli.NewApp()
	app.Name = "XChange"
	app.Version = "v0.1"
	app.Description = "Application for currency rates collection"
	app.Usage = ""

	var debugLogging bool
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
	}

	app.Before = func(c *cli.Context) error {
		logger.Get().Formatter = &logrus.TextFormatter{FullTimestamp: true}
		if c.Bool("debug") {
			logger.SetDebugLevel()
			logger.Get().Debug("Debug logging enabled")
			logger.Get().Debug(app.Name, "-", app.Version)
		}
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:  "collect",
			Usage: "Collects currency rates",
			Action: func(ctx *cli.Context) error {
				return command.CollectExecute(ctx.Context, ctx.Bool("dryRun"))
			},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name: "dryRun, r",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Get().Fatal(err)
	}
}
