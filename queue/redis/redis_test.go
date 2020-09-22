// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"fmt"
	"testing"

	"github.com/Bose/minisentinel"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

// setup global variables used for testing
var (
	_build = &library.Build{
		ID:           vela.Int64(1),
		Number:       vela.Int(1),
		Parent:       vela.Int(1),
		Event:        vela.String("push"),
		Status:       vela.String("success"),
		Error:        vela.String(""),
		Enqueued:     vela.Int64(1563474077),
		Created:      vela.Int64(1563474076),
		Started:      vela.Int64(1563474077),
		Finished:     vela.Int64(0),
		Deploy:       vela.String(""),
		Clone:        vela.String("https://github.com/github/octocat.git"),
		Source:       vela.String("https://github.com/github/octocat/abcdefghi123456789"),
		Title:        vela.String("push received from https://github.com/github/octocat"),
		Message:      vela.String("First commit..."),
		Commit:       vela.String("48afb5bdc41ad69bf22588491333f7cf71135163"),
		Sender:       vela.String("OctoKitty"),
		Author:       vela.String("OctoKitty"),
		Branch:       vela.String("master"),
		Ref:          vela.String("refs/heads/master"),
		BaseRef:      vela.String(""),
		Host:         vela.String("example.company.com"),
		Runtime:      vela.String("docker"),
		Distribution: vela.String("linux"),
	}

	_repo = &library.Repo{
		ID:          vela.Int64(1),
		Org:         vela.String("github"),
		Name:        vela.String("octocat"),
		FullName:    vela.String("github/octocat"),
		Link:        vela.String("https://github.com/github/octocat"),
		Clone:       vela.String("https://github.com/github/octocat.git"),
		Branch:      vela.String("master"),
		Timeout:     vela.Int64(60),
		Visibility:  vela.String("public"),
		Private:     vela.Bool(false),
		Trusted:     vela.Bool(false),
		Active:      vela.Bool(true),
		AllowPull:   vela.Bool(false),
		AllowPush:   vela.Bool(true),
		AllowDeploy: vela.Bool(false),
		AllowTag:    vela.Bool(false),
	}

	_steps = &pipeline.Build{
		Version: "1",
		ID:      "github_octocat_1",
		Services: pipeline.ContainerSlice{
			{
				ID:          "service_github_octocat_1_postgres",
				Directory:   "/home/github/octocat",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "postgres:12-alpine",
				Name:        "postgres",
				Number:      1,
				Ports:       []string{"5432:5432"},
				Pull:        "not_present",
			},
		},
		Steps: pipeline.ContainerSlice{
			{
				ID:          "step_github_octocat_1_init",
				Directory:   "/home/github/octocat",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "#init",
				Name:        "init",
				Number:      1,
				Pull:        "always",
			},
			{
				ID:          "step_github_octocat_1_clone",
				Directory:   "/home/github/octocat",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "target/vela-git:v0.3.0",
				Name:        "clone",
				Number:      2,
				Pull:        "always",
			},
			{
				ID:          "step_github_octocat_1_echo",
				Commands:    []string{"echo hello"},
				Directory:   "/home/github/octocat",
				Environment: map[string]string{"FOO": "bar"},
				Image:       "alpine:latest",
				Name:        "echo",
				Number:      3,
				Pull:        "always",
			},
		},
	}

	_user = &library.User{
		ID:     vela.Int64(1),
		Name:   vela.String("octocat"),
		Token:  vela.String("superSecretToken"),
		Hash:   nil,
		Active: vela.Bool(true),
		Admin:  vela.Bool(false),
	}
)

func TestRedis_New_Success(t *testing.T) {
	// setup redis
	redis, _ := miniredis.Run()

	// setup types
	uri := fmt.Sprintf("redis://%s", redis.Addr())

	// run test
	_, err := New(uri, constants.DefaultRoute)
	if err != nil {
		t.Error("New should not have returned err: ", err)
	}
}

func TestRedis_New_Failure(t *testing.T) {
	// setup redis
	redis, _ := miniredis.Run()

	tests := []struct {
		data string
		want error
	}{
		{ // connection uri with invalid URI
			data: redis.Addr(),
			want: fmt.Errorf(""),
		},
	}

	// run tests
	for _, test := range tests {
		// run test
		_, err := New(test.data, constants.DefaultRoute)
		if err == nil {
			t.Errorf("New should have returned err")
		}
	}
}

func TestRedis_NewCluster_Success(t *testing.T) {
	// setup redis
	replica, _ := miniredis.Run()
	redis := minisentinel.NewSentinel(replica, minisentinel.WithReplica(replica))
	_ = redis.Start()

	// setup types
	uri := fmt.Sprintf("redis://%s,%s", redis.MasterInfo().Name, redis.Addr())

	// run test
	_, err := NewCluster(uri, constants.DefaultRoute)
	if err != nil {
		t.Error("NewCluster should not have returned err: ", err)
	}
}

func TestRedis_NewCluster_Failure(t *testing.T) {
	// setup redis
	replica, _ := miniredis.Run()
	redis := minisentinel.NewSentinel(replica, minisentinel.WithReplica(replica))
	_ = redis.Start()

	tests := []struct {
		data string
		want error
	}{
		{ // connection uri with invalid URI
			data: fmt.Sprintf("%s,%s", redis.MasterInfo().Name, redis.Addr()),
			want: fmt.Errorf(""),
		},
		{ // connection uri that will timeout
			data: fmt.Sprintf("redis://%s,%s", redis.MasterInfo().Name, "localhost"),
			want: fmt.Errorf(""),
		},
	}

	// run tests
	for _, test := range tests {
		// run test
		_, err := New(test.data, constants.DefaultRoute)
		if err == nil {
			t.Errorf("New should have returned err")
		}
	}
}
