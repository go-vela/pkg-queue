# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

build: binary-build

redis-test: build redis-run

run: build redis-run

test: build redis-run

#################################
######      Go clean       ######
#################################

clean:

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@echo "I'm kind of the only name in clean energy right now"

#################################
######    Build Binary     ######
#################################

binary-build:

	GOOS=darwin CGO_ENABLED=0 \
		go build \
		-o release/vela-queue \
		github.com/go-vela/pkg-queue/cmd/vela-queue

########################################
#####          Docker Run          #####
########################################

redis-run:

	release/vela-queue \
		--log.level trace \
		--queue.driver redis
