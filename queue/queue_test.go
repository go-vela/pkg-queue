// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package queue

import (
	"fmt"
	"strings"
	"testing"
)

func TestQueue_New_Success(t *testing.T) {

	//TODO: come back to this once we have a good mock pattern in place.
	// // setup types
	// s := &Setup{Driver: constants.DriverRedis, Config: "redis://redis"}
	// redis, _ := s.Redis()

	// tests := []struct {
	// 	data *Setup
	// 	want Service
	// }{
	// 	{ // test for Redis Queue
	// 		data: &Setup{Driver: constants.DriverRedis, Config: "redis://redis"},
	// 		want: redis,
	// 	},
	// }

	// // run tests
	// for _, test := range tests {
	// 	// run test
	// 	_, err := New(test.data)
	// 	if err != nil {
	// 		t.Errorf("New should not have returned err: %w", err)
	// 	}
	// }
}

func TestQueue_New_Failure(t *testing.T) {
	tests := []struct {
		data *Setup
		want error
	}{
		{ // test for unsupported kafka
			data: &Setup{Driver: "kafka", Config: "localhost:9946"},
			want: fmt.Errorf("unsupported queue driver: kafka"),
		},
		{ // test for invalid queues
			data: &Setup{Driver: "foobar", Config: "bad:config"},
			want: fmt.Errorf("invalid queue driver: foobar"),
		},
	}

	// run tests
	for _, test := range tests {
		// run test
		_, err := New(test.data)
		if err == nil {
			t.Error("New should have returned err")
		}

		if !strings.EqualFold(err.Error(), test.want.Error()) {
			t.Errorf("Err is %v, want %v", err, test.want)
		}
	}
}
