// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/go-vela/types"
	"github.com/sirupsen/logrus"
)

// Pop grabs an item from the specified channel off the queue.
func (c *client) Pop() (*types.Item, error) {
	logrus.Tracef("popping item from queue %s", c.Channels)

	// build a redis queue command to pop an item from queue
	//
	// https://pkg.go.dev/github.com/go-redis/redis?tab=doc#Client.BLPop
	popCmd := c.Queue.BLPop(context.Background(), c.Timeout, c.Channels...)

	// blocking call to pop item from queue
	//
	// https://pkg.go.dev/github.com/go-redis/redis?tab=doc#StringSliceCmd.Result
	result, err := popCmd.Result()
	if err != nil {
		switch err {
		case redis.Nil: // BLPOP timeout
			return nil, nil
		default:
			return nil, err
		}
	}

	item := new(types.Item)

	// unmarshal result into queue item
	err = json.Unmarshal([]byte(result[1]), item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
