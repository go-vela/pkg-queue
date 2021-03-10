// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"fmt"
	"strings"

	"github.com/go-vela/pkg-queue/queue/redis"
	"github.com/go-vela/types/constants"
	"github.com/sirupsen/logrus"
)

// Setup represents the configuration necessary for
// creating a Vela service capable of integrating
// with a configured queue environment.
type Setup struct {
	// specifies the queue driver to use
	Driver string
	// enables the queue client to integrate with a cluster
	Cluster bool
	// configuration string for the queue
	Config string
	// channels to listen on for the queue
	Routes []string
}

// Redis creates and returns a Vela engine capable of
// integrating with a Redis queue.
func (s *Setup) Redis() (Service, error) {
	// check if the default route is provided
	if !strings.Contains(strings.Join(s.Routes, ","), constants.DefaultRoute) {
		s.Routes = append(s.Routes, constants.DefaultRoute)
	}

	// create new Redis queue service
	//
	// https://pkg.go.dev/github.com/go-vela/pkg-queue/queue/redis?tab=doc#New
	return redis.New(
		redis.WithAddress(s.Config),
		redis.WithChannels(s.Routes...),
		redis.WithCluster(s.Cluster),
	)
}

// Kafka creates and returns a Vela engine capable of
// integrating with a Kafka queue.
func (s *Setup) Kafka() (Service, error) {
	logrus.Tracef("Creating %s queue client from CLI configuration", constants.DriverKafka)
	// return kafka.New(c.String("queue-config"), "vela")
	return nil, fmt.Errorf("unsupported queue driver: %s", constants.DriverKafka)
}

// Validate verifies the necessary fields for the
// provided configuration are populated correctly.
func (s *Setup) Validate() error {
	logrus.Trace("Validating queue CLI configuration")

	if len(s.Driver) == 0 {
		return fmt.Errorf("queue.driver (VELA_QUEUE_DRIVER or QUEUE_DRIVER) flag not specified")
	}

	if len(s.Config) == 0 {
		return fmt.Errorf("queue.config (VELA_QUEUE_CONFIG or QUEUE_CONFIG) flag not specified")
	}

	// setup is valid
	return nil
}
