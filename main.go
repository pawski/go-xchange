package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"github.com/urfave/cli"
	"github.com/pawski/go-xchange/command"
	"github.com/pawski/go-xchange/logger"
)

func main() {

	app := cli.NewApp()
	app.Name = "XChange"
	app.Version = "v0.1"
	app.Description = "Application for currency rates collection"
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
		},{
			Name:  "fetch",
			Usage: "Fetch currency rates",
			Action: func(c *cli.Context) {
				if err := command.FetchExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		},{
			Name: "balance",
			Usage: "Fetch Account balance",
			Action: func(c *cli.Context) {
				if err := command.BalanceExecute(); err != nil {
					logger.Get().Error(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
