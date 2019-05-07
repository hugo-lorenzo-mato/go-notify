package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hugo-lorenzo-mato/go-notify/cmd"
	"github.com/hugo-lorenzo-mato/go-notify/conf"
	"github.com/urfave/cli"
	"os"
)

func init() {
	//If environment variable DEBUG is present we change to debug log level
	log.SetFormatter(&log.TextFormatter{})
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
		log.Debug("Log level set to DEBUG")
	}
}

func setDebug() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Debug("Log level set to DEBUG")
}

func main() {

	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Println(`   

   ____  ___        _   _  ___ _____ ___ _______   __
  / ___|/ _ \      | \ | |/ _ \_   _|_ _|  ___\ \ / /
 | |  _| | | |_____|  \| | | | || |  | || |_   \ V / 
 | |_| | |_| |_____| |\  | |_| || |  | ||  _|   | |  
  \____|\___/      |_| \_|\___/ |_| |___|_|     |_|
			           (` + "Version " + conf.Version + `)
`)
	}

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:        "debug, d",
			Usage:       "Activate debug mode",
			Destination: &conf.Debug,
		},
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Yaml config file path",
			Destination: &conf.YamlPath,
			Value:       ".",
		},
	}

	if conf.Debug {
		setDebug()
	}
	app.Version = conf.Version
	app.Usage = "Reports any type of incident to the configured messaging services"
	app.Commands = cmd.Cmds

	app.Run(os.Args)

}
