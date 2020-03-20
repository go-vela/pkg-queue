// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"fmt"

	"github.com/go-vela/types/constants"
	"github.com/sirupsen/logrus"
)

// New creates and returns a Vela service capable of integrating
// with the configured queue environments. Currently the
// following queues are supported:
//
// * redis
func New(s *Setup) (Service, error) {
	// validate the setup being provided
	err := s.Validate()
	if err != nil {
		return nil, err
	}

	logrus.Debug("creating queue client from setup")
	// process the queue driver being provided
	switch s.Driver {
	case constants.DriverKafka:
		return s.Kafka()
	case constants.DriverRedis:
		return s.Redis()
	default:
		return nil, fmt.Errorf("invalid queue driver: %s", s.Driver)
	}
}
