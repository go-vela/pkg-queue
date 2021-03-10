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

func TestQueue_New(t *testing.T) {
	// setup types

	// create a local fake redis instance
	//
	// https://pkg.go.dev/github.com/alicebob/miniredis/v2#Run
	mr, err := miniredis.Run()
	if err != nil {
		t.Errorf("unable to create miniredis instance: %v", err)
	}

	_redis, err := redis.New(
		redis.WithAddress(mr.Addr()),
		redis.WithChannels("foo"),
		redis.WithCluster(false),
	)
	if err != nil {
		t.Errorf("unable to create redis service: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		setup   *Setup
		want    Service
	}{
		{
			failure: false,
			setup: &Setup{
				Driver:  "redis",
				Address: mr.Addr(),
				Routes:  []string{"foo"},
				Cluster: false,
			},
			want: _redis,
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "kafka",
				Address: "kafka://kafka.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
			want: nil,
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "pubsub",
				Address: "pubsub://pubsub.example.com",
				Routes:  []string{"foo"},
				Cluster: false,
			},
			want: nil,
		},
		{
			failure: true,
			setup: &Setup{
				Driver:  "redis",
				Address: "",
				Routes:  []string{"foo"},
				Cluster: false,
			},
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		got, err := New(test.setup)

		if test.failure {
			if err == nil {
				t.Errorf("New should have returned err")
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("New is %v, want %v", got, test.want)
			}

			continue
		}

		if err != nil {
			t.Errorf("New returned err: %v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("New is %v, want %v", got, test.want)
		}
	}
}
