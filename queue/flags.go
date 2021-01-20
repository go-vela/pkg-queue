// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import "github.com/urfave/cli/v2"

// Flags represents all supported command line
// interface (CLI) flags for the queue.
//
// https://pkg.go.dev/github.com/urfave/cli?tab=doc#Flag
var Flags = []cli.Flag{

	&cli.StringFlag{
		EnvVars: []string{"QUEUE_LOG_LEVEL", "VELA_LOG_LEVEL", "LOG_LEVEL"},
		Name:    "queue.log.level",
		Usage:   "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
		Value:   "info",
	},

	// Queue Flags

	&cli.StringFlag{
		EnvVars: []string{"VELA_QUEUE_DRIVER", "QUEUE_DRIVER"},
		Name:    "queue.driver",
		Usage:   "queue driver",
	},
	&cli.StringFlag{
		EnvVars: []string{"VELA_QUEUE_CONFIG", "QUEUE_CONFIG"},
		Name:    "queue.config",
		Usage:   "queue driver configuration string",
	},
	&cli.BoolFlag{
		EnvVars: []string{"VELA_QUEUE_CLUSTER", "QUEUE_CLUSTER"},
		Name:    "queue.cluster",
		Usage:   "queue client is setup for clusters",
	},
	// By default all builds are pushed to the "vela" route
	&cli.StringSliceFlag{
		EnvVars: []string{"VELA_QUEUE_WORKER_ROUTES", "QUEUE_WORKER_ROUTES"},
		Name:    "queue.worker.routes",
		Usage:   "queue worker routes is configuration for routing builds",
	},
}
