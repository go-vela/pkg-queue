// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"reflect"
	"testing"

	"github.com/go-vela/types"
	"gopkg.in/square/go-jose.v2/json"
)

func TestRedis_Pop_Success(t *testing.T) {
	// setup redis mock
	c, _ := NewTest("vela")

	// set types
	//
	// use global variables in redis_test.go
	want := &types.Item{
		Build:    _build,
		Pipeline: _steps,
		Repo:     _repo,
		User:     _user,
	}

	// seed queue
	item, _ := json.Marshal(want)

	err := c.Queue.RPush("vela", item).Err()
	if err != nil {
		t.Error("RPush should not have returned err: ", err)
	}

	// run test
	got, err := c.Pop()
	if err != nil {
		t.Error("Pop should not have returned err: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pop is %v, want %v", got, want)
	}
}

func TestRedis_Pop_BadChannel(t *testing.T) {
	// setup redis mock
	c, _ := NewTest("vela")

	// overwrite channel to be invalid
	c.Channels = nil

	err := c.Queue.RPush("vela", nil).Err()
	if err != nil {
		t.Error("RPush should not have returned err: ", err)
	}

	// run test
	_, err = c.Pop()
	if err == nil {
		t.Error("Pop should have returned err")
	}
}

func TestRedis_Pop_BadItem(t *testing.T) {
	// setup redis mock
	c, _ := NewTest("vela")

	err := c.Queue.RPush("vela", nil).Err()
	if err != nil {
		t.Error("RPush should not have returned err: ", err)
	}

	// run test
	_, err = c.Pop()
	if err == nil {
		t.Error("Pop should have returned err")
	}
}
