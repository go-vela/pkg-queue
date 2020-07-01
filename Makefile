# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

.PHONY: run
run: build redis-run

.PHONY: test
test: redis-test

.PHONY: redis-test
redis-test: build redis-run

.PHONY: build
build: binary-build

#################################
######      Go clean       ######
#################################

.PHONY: clean
clean:

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@echo "I'm kind of the only name in clean energy right now"

#################################
######    Build Binary     ######
#################################

.PHONY: binary-build
binary-build:

	GOOS=darwin CGO_ENABLED=0 \
		go build \
		-o release/vela-queue \
		github.com/go-vela/pkg-queue/cmd/vela-queue

########################################
#####          Docker Run          #####
########################################

.PHONY: redis-run
redis-run:

	release/vela-queue \
		--log.level trace \
		--queue.driver redis
