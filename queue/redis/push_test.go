// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"encoding/json"
	"testing"

	"github.com/go-vela/types"
)

func TestRedis_Push_Success(t *testing.T) {
	// setup redis mock
	c, _ := NewTest("vela")

	// set types
	//
	// use global variables in redis_test.go
	item, _ := json.Marshal(&types.Item{
		Build:    _build,
		Pipeline: _steps,
		Repo:     _repo,
		User:     _user,
	})

	// run test
	err := c.Push("vela", item)
	if err != nil {
		t.Error("Pop should not have returned err: ", err)
	}
}
