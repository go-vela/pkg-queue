// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Package Information

	app.Name = "vela-queue"
	app.HelpName = "vela-queue"
	app.Usage = "Vela queue package for integrating with different queues"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []cli.Author{
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

		cli.StringFlag{
			EnvVar: "VELA_LOG_LEVEL,QUEUE_LOG_LEVEL",
			Name:   "log.level",
			Usage:  "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:  "info",
		},

		// Queue Flags

		cli.StringFlag{
			EnvVar: "VELA_QUEUE_DRIVER,QUEUE_DRIVER",
			Name:   "queue.driver",
			Usage:  "queue driver",
		},
		cli.StringFlag{
			EnvVar: "VELA_QUEUE_CONFIG,QUEUE_CONFIG",
			Name:   "queue.config",
			Usage:  "queue driver configuration string",
		},
		cli.BoolFlag{
			EnvVar: "VELA_QUEUE_CLUSTER,QUEUE_CLUSTER",
			Name:   "queue.cluster",
			Usage:  "queue client is setup for clusters",
		},
		// By default all builds are pushed to the "vela" route
		cli.StringSliceFlag{
			EnvVar: "VELA_QUEUE_WORKER_ROUTES,QUEUE_WORKER_ROUTES",
			Name:   "queue.worker.routes",
			Usage:  "queue worker routes is configuration for routing builds",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
