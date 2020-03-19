// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"github.com/go-vela/types"
	"github.com/go-vela/types/pipeline"
)

// Service represents the interface for Vela integrating
// with the different supported Queue backends.
type Service interface {
	// Pop defines a function that grabs an
	// item off the queue.
	Pop() (*types.Item, error)

	// Push defines a function that publishes an
	// item to the specified route in the queue.
	Push(string, []byte) error

	// Publish defines a function that decides which
	// route a build gets placed within the queue.
	Route(*pipeline.Worker) (string, error)
}
