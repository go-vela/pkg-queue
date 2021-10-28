# pkg-queue

> DISCLAIMER:
>
> The contents of this repository have been migrated into [go-vela/server](https://github.com/go-vela/server).
>
> This was done as a part of [go-vela/community#394](https://github.com/go-vela/community/issues/394) to deliver [on a proposal](https://github.com/go-vela/community/blob/master/proposals/2021/08-25_repo-structure.md).

[![license](https://img.shields.io/crates/l/gl.svg)](../LICENSE)
[![GoDoc](https://godoc.org/github.com/go-vela/pkg-queue?status.svg)](https://godoc.org/github.com/go-vela/pkg-queue)
[![Go Report Card](https://goreportcard.com/badge/go-vela/pkg-queue)](https://goreportcard.com/report/go-vela/pkg-queue)
[![codecov](https://codecov.io/gh/go-vela/pkg-queue/branch/master/graph/badge.svg)](https://codecov.io/gh/go-vela/pkg-queue)

Vela package is designed to publish build items in work queues between [go-vela/server](https://github.com/go-vela/server) and [go-vela/worker](https://github.com/go-vela/worker).

The following queues are supported:

* [Redis](https://docker.io/)

## Documentation

For installation and usage, please [visit our docs](https://go-vela.github.io/docs).

## Contributing

We are always welcome to new pull requests!

Please see our [contributing](CONTRIBUTING.md) docs for further instructions.

## Support

We are always here to help!

Please see our [support](SUPPORT.md) documentation for further instructions.

## Copyright and License

```
Copyright (c) 2021 Target Brands, Inc.
```

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)
