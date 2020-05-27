package main

import (
	"github.com/sirupsen/logrus"
	"github.com/pawski/go-xchange/application/command"
	"github.com/pawski/go-xchange/logger"
	"github.com/urfave/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "XChange"
	app.Version = "v0.1"
	app.Description = "Application for currency exchange support"
	app.Usage = ""

	var debugLogging bool
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug logging",
			Destination: &debugLogging,
		},
	}

	app.Before = func(c *cli.Context) error {
		logger.Get().Formatter = &logrus.TextFormatter{FullTimestamp: true}
		if c.GlobalBool("debug") {
			logger.SetDebugLevel()
			logger.Get().Debug("Debug logging enabled")
			logger.Get().Debug(app.Name, "-", app.Version)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "collect",
			Usage: "Collects currency rates",
			Action: func(c *cli.Context) {
				if err := command.CollectExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		}, {
			Name:  "fetch",
			Usage: "Fetch currency rates",
			Action: func(c *cli.Context) {
				if err := command.FetchExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		}, {
			Name:  "balance",
			Usage: "Fetch Account balance",
			Action: func(c *cli.Context) {
				if err := command.BalanceExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		}, {
			Name:  "directrates",
			Usage: "Fetch Direct Rates",
			Action: func(c *cli.Context) {
				if err := command.DirectRatesExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		}, {
			Name:  "run",
			Usage: "Runs continuous monitoring",
			Action: func(c *cli.Context) {
				command.RunExecute()
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logger.Get().Fatal(err)
	}
}
