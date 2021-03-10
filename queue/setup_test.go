// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"reflect"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-vela/pkg-queue/queue/redis"
)

func TestQueue_Setup_Redis(t *testing.T) {
	// setup types

	// create a local fake redis instance
	//
	// https://pkg.go.dev/github.com/alicebob/miniredis/v2#Run
	mr, err := miniredis.Run()
	if err != nil {
		t.Errorf("unable to create miniredis instance: %v", err)
	}

	_setup := &Setup{
		Driver:  "redis",
		Address: mr.Addr(),
		Routes:  []string{"foo"},
		Cluster: false,
	}

	want, err := redis.New(
		redis.WithAddress(mr.Addr()),
		redis.WithChannels("foo"),
		redis.WithCluster(false),
	)
	if err != nil {
		t.Errorf("unable to create redis service: %v", err)
	}

	got, err := _setup.Redis()
	if err != nil {
		t.Errorf("Redis returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Redis is %v, want %v", got, want)
	}
}

func TestQueue_Setup_Kafka(t *testing.T) {
	// setup types
	_setup := &Setup{
		Driver:  "kafka",
		Address: "kafka://kafka.example.com",
		Routes:  []string{"foo"},
		Cluster: false,
	}

	got, err := _setup.Kafka()
	if err == nil {
		t.Errorf("Kafka should have returned err")
	}

	if got != nil {
		t.Errorf("Kafka is %v, want nil", got)
	}
}

func TestSource_Setup_Validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		setup   *Setup
	}{
		{
			failure: false,
			setup: &Setup{
				Driver:  "redis",
				Address: "redis://redis.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: false,
			setup: &Setup{
				Driver:  "kafka",
				Address: "kafka://kafka.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "redis",
				Address: "redis://redis.example.com/",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "redis",
				Address: "redis.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "",
				Address: "redis://redis.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "redis",
				Address: "",
				Routes:  []string{"foo"},
				Cluster: false,
			},
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "redis",
				Address: "redis://redis.example.com",
				Routes:  []string{},
				Cluster: false,
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.setup.Validate()

		if test.failure {
			if err == nil {
				t.Errorf("Validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Validate returned err: %v", err)
		}
	}
}
