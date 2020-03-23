// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/Bose/minisentinel"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-vela/types/constants"
)

func TestQueue_Kafka(t *testing.T) {
	// setup types
	s := &Setup{Driver: constants.DriverKafka, Config: "localhost:9946"}
	want := fmt.Errorf("unsupported queue driver: %s", constants.DriverKafka)

	// run test
	got, err := s.Kafka()
	if err == nil {
		t.Error("Kafka should have returned err")
	}

	if !(got == nil) {
		t.Errorf("Kafka is %+v, want %+v", got, nil)
	}

	if !reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(want)) {
		t.Errorf("Kafka is %+v, want %+v", got, want)
	}
}

func TestQueue_Redis(t *testing.T) {
	// setup redis
	replica, _ := miniredis.Run()
	redis := minisentinel.NewSentinel(replica, minisentinel.WithReplica(replica))
	_ = redis.Start()

	tests := []struct {
		data *Setup
		want error
	}{
		{ // test non for clustered redis client
			data: &Setup{
				Driver:  constants.DriverRedis,
				Config:  fmt.Sprintf("redis://%s", replica.Addr()),
				Cluster: false,
				Routes:  []string{},
			},
			want: nil,
		},
		{ // test non for cluster redis client
			data: &Setup{
				Driver:  constants.DriverRedis,
				Config:  fmt.Sprintf("redis://%s,%s", redis.MasterInfo().Name, redis.Addr()),
				Cluster: true,
				Routes:  []string{},
			},
			want: nil,
		},
	}

	// run tests
	for _, test := range tests {
		// run test
		_, err := test.data.Redis()
		if err != nil {
			t.Error("Redis should not have returned err: ", err)
		}
	}
}

func TestQueue_Validate(t *testing.T) {
	// setup types
	tests := []struct {
		data *Setup
		want error
	}{
		{
			// test if the queue setup is empty
			data: &Setup{},
			want: fmt.Errorf("queue.driver (VELA_QUEUE_DRIVER or QUEUE_DRIVER) flag not specified"),
		},
		{
			// test if the queue provided is set with default value
			data: &Setup{Driver: "foobar"},
			want: fmt.Errorf("queue.config (VELA_QUEUE_CONFIG or QUEUE_CONFIG) flag not specified"),
		},
		{
			// test if the queue provided is set with default value
			data: &Setup{Config: ""},
			want: fmt.Errorf("queue.driver (VELA_QUEUE_DRIVER or QUEUE_DRIVER) flag not specified"),
		},
	}

	// run tests
	for _, test := range tests {
		// run test
		err := test.data.Validate()
		if err == nil {
			t.Error("Validate should have returned err")
		}

		if !strings.EqualFold(err.Error(), test.want.Error()) {
			t.Errorf("Err is %v, want %v", err, test.want)
		}
	}
}
