// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import "github.com/urfave/cli/v2"

// Flags represents all supported command line
// interface (CLI) flags for the queue.
var Flags = []cli.Flag{
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
