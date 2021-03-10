// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"reflect"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func TestRedis_ClientOpt_WithAddress(t *testing.T) {
	// setup tests

	// create a local fake redis instance
	//
	// https://pkg.go.dev/github.com/alicebob/miniredis/v2#Run
	mr, err := miniredis.Run()
	if err != nil {
		t.Errorf("unable to create miniredis instance: %v", err)
	}

	tests := []struct {
		failure bool
		address string
		want    string
	}{
		{
			failure: false,
			address: mr.Addr(),
			want:    mr.Addr(),
		},
		{
			failure: true,
			address: "",
			want:    "",
		},
	}

	// run tests
	for _, test := range tests {
		_service, err := New(
			WithAddress(test.address),
		)

		if test.failure {
			if err == nil {
				t.Errorf("WithAddress should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("WithAddress returned err: %v", err)
		}

		if !reflect.DeepEqual(_service.config.Address, test.want) {
			t.Errorf("WithAddress is %v, want %v", _service.config.Address, test.want)
		}
	}
}

func TestRedis_ClientOpt_WithChannels(t *testing.T) {
	// setup tests
	tests := []struct {
		failure  bool
		channels []string
		want     []string
	}{
		{
			failure:  false,
			channels: []string{"foo", "bar"},
			want:     []string{"foo", "bar"},
		},
		{
			failure:  true,
			channels: []string{},
			want:     []string{},
		},
	}

	// run tests
	for _, test := range tests {
		_service, err := New(
			WithChannels(test.channels...),
		)

		if test.failure {
			if err == nil {
				t.Errorf("WithChannels should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("WithChannels returned err: %v", err)
		}

		if !reflect.DeepEqual(_service.config.Channels, test.want) {
			t.Errorf("WithChannels is %v, want %v", _service.config.Channels, test.want)
		}
	}
}

func TestRedis_ClientOpt_WithCluster(t *testing.T) {
	// setup tests
	tests := []struct {
		cluster bool
		want    bool
	}{
		{
			cluster: true,
			want:    true,
		},
		{
			cluster: false,
			want:    false,
		},
	}

	// run tests
	for _, test := range tests {
		_service, err := New(
			WithCluster(test.cluster),
		)

		if err != nil {
			t.Errorf("WithCluster returned err: %v", err)
		}

		if !reflect.DeepEqual(_service.config.Cluster, test.want) {
			t.Errorf("WithCluster is %v, want %v", _service.config.Cluster, test.want)
		}
	}
}
