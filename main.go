package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"github.com/urfave/cli"
	"github.com/pawski/go-xchange/command"
)

func main() {

	logger := logrus.New()

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
		logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
		if c.GlobalBool("debug") {
			logger.Level = logrus.DebugLevel
			logger.Debug("Debug logging enabled")
			logger.Debug(app.Name, "-", app.Version)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "collect",
			Usage: "Starts continues collection of currency rates",
			Action: func(c *cli.Context) {
				if err := command.CollectExecute(); err != nil {
					logger.Error(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
