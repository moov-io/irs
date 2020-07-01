moov-io/irs
===

[![GoDoc](https://godoc.org/github.com/moov-io/irs?status.svg)](https://godoc.org/github.com/moov-io/irs)
[![Build Status](https://travis-ci.com/moov-io/irs.svg?branch=master)](https://travis-ci.com/moov-io/irs)
[![Coverage Status](https://codecov.io/gh/moov-io/irs/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/irs)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/irs)](https://goreportcard.com/report/github.com/moov-io/irs)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/irs/master/LICENSE)

Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and 
CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns 
Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and 
CSV/TEXT file formats.

| Input      | Output     |
|------------|------------|
| JSON       | JSON       |
| ASCII FIRE | ASCII FIRE |
|            | PDF Form   |
|            | SQL        |


Docs: [docs](docs/README.md) | [open api specification](api/api.yml)

## Project Status

We are just getting started! 

- [ ] 1099-MISC [About Form 1099-MISC](https://www.irs.gov/forms-pubs/about-form-1099-misc)
- [ ] 1099-NEC [About Form 1099-NEC](https://www.irs.gov/forms-pubs/about-form-1099-nec)  

... more to come 

## Getting Started

Read through the [project docs](docs/README.md) over here to get an understanding of the purpose of this project and how to run it. 

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/irs/contribute) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/irs/releases/latest) as well. We highly recommend you use a tagged release for production.

### Test Coverage

Improving test coverage is a good candidate for new contributors while also allowing the project to move more quickly by reducing regressions issues that might not be caught before a release is pushed out to our users. One great way to improve coverage is by adding edge cases and different inputs to functions (or [contributing and running fuzzers](https://github.com/dvyukov/go-fuzz)).

Tests can run processes (like sqlite databases), but should only do so locally.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
