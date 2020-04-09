// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"time"

	"github.com/go-vela/pkg-queue/queue"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Package Information

	app.Name = "vela-queue"
	app.HelpName = "vela-queue"
	app.Usage = "Vela queue package for integrating with different queues"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Package Metadata

	app.Compiled = time.Now()
	app.Action = run

	// Package Flags

	app.Flags = []cli.Flag{

		&cli.StringFlag{
			EnvVars: []string{"VELA_LOG_LEVEL", "QUEUE_LOG_LEVEL"},
			Name:    "log.level",
			Usage:   "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:   "info",
		},
	}

	// Queue Flags
	app.Flags = append(app.Flags, queue.Flags...)

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
