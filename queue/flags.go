// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"time"

  "github.com/go-vela/types/constants"
	"github.com/urfave/cli/v2"
)

// Flags represents all supported command line
// interface (CLI) flags for the queue.
//
// https://pkg.go.dev/github.com/urfave/cli?tab=doc#Flag
var Flags = []cli.Flag{

	// Logger Flags

	&cli.StringFlag{
		EnvVars: []string{"QUEUE_LOG_FORMAT", "VELA_LOG_FORMAT", "LOG_FORMAT"},
		Name:    "queue.log.format",
		Usage:   "format of logs to output",
		Value:   "json",
	},
	&cli.StringFlag{
		EnvVars: []string{"QUEUE_LOG_LEVEL", "VELA_LOG_LEVEL", "LOG_LEVEL"},
		Name:    "queue.log.level",
		Usage:   "level of logs to output",
		Value:   "info",
	},

	// Queue Flags

	&cli.StringFlag{
		EnvVars: []string{"VELA_QUEUE_DRIVER", "QUEUE_DRIVER"},
		Name:    "queue.driver",
		Usage:   "driver to be used for the queue",
	},
	&cli.StringFlag{
		EnvVars: []string{"VELA_QUEUE_ADDR", "QUEUE_ADDR"},
		Name:    "queue.addr",
		Usage:   "fully qualified url (<scheme>://<host>) for the queue",
	},
	&cli.BoolFlag{
		EnvVars: []string{"VELA_QUEUE_CLUSTER", "QUEUE_CLUSTER"},
		Name:    "queue.cluster",
		Usage:   "enables connecting to a queue cluster",
	},
	&cli.StringSliceFlag{
		EnvVars: []string{"VELA_QUEUE_ROUTES", "QUEUE_ROUTES"},
		Name:    "queue.routes",
		Usage:   "list of routes (channels/topics) to publish builds",
		Value:   cli.NewStringSlice(constants.DefaultRoute),
	},
	&cli.DurationFlag{
		EnvVars: []string{"VELA_QUEUE_BLPOP_TIMEOUT", "QUEUE_BLPOP_TIMEOUT"},
		Name:    "queue.worker.blpop.timeout",
		Usage:   "queue timeout for the blpop call",
		Value:   60 * time.Second,
	},
}
