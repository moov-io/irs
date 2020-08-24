moov-io/irs
===

[![GoDoc](https://godoc.org/github.com/moov-io/irs?status.svg)](https://godoc.org/github.com/moov-io/irs)
[![Build Status](https://travis-ci.com/moov-io/irs.svg?branch=master)](https://travis-ci.com/moov-io/irs)
[![Coverage Status](https://codecov.io/gh/moov-io/irs/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/irs)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/irs)](https://goreportcard.com/report/github.com/moov-io/irs)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/irs/master/LICENSE)

Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and
CLI for creating, parsing, validating, and transforming IRS electronic [Filing Information Returns
Electronically](https://www.irs.gov/e-file-providers/filing-information-returns-electronically-fire) (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and
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

## Commands

Irs has command line interface to manage irs files and to lunch web service.

```
irs --help
```
```
Usage:
   [command]

Available Commands:
  convert     Convert irs file format
  help        Help about any command
  print       Print irs file
  validator   Validate irs file
  web         Launches web server

Flags:
  -h, --help           help for this command
      --input string   input file (default is $PWD/irs.json)

Use " [command] --help" for more information about a command.
```

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert from a irs file to another format file. Result will create a irs file.
`print` | The print command allows users to print a irs file with special file format (json, irs).
`validator` | The validator command allows users to validate a irs file.
`web` | The web command will launch a web server with endpoints to manage irs files.

### file convert

```
irs convert --help
```
```
Usage:
   convert [output] [flags]

Flags:
      --format string   format of irs file(required) (default "json")
  -h, --help            help for convert

Global Flags:
      --input string   input file (default is $PWD/irs.json)
```

The output parameter is the full path name to convert new irs file.
The format parameter is supported 2 types, "json" and  "irs".
The generate parameter will replace new generated trailer record in the file.
The input parameter is source irs file, supported raw type file and json type file.

example:
```
irs convert output/output.json --input testdata/packed_file.json --format json
```

### file print

```
irs print --help
```
```
Usage:
   print [flags]

Flags:
      --format string   print format (default "json")
  -h, --help            help for print

Global Flags:
      --input string   input file (default is $PWD/irs.json)
```

The format parameter is supported 2 types, "json" and  "irs".
The input parameter is source irs file, supported raw type file and json type file.

### file validate

```
irs validator --help
```
```
Usage:
   validator [flags]

Flags:
  -h, --help   help for validator

Global Flags:
      --input string   input file (default is $PWD/irs.json)
```

The input parameter is source irs file, supported raw type file and json type file.

example:
```
irs validator --input testdata/packed_file.dat
Error: is an invalid value of TotalConsumerSegmentsJ1

irs validator --input testdata/packed_file.json
```

### web server

```
irs web --help
```
```
Usage:
   web [flags]

Flags:
  -h, --help          help for web
  -t, --test          test server

Global Flags:
      --input string   input file (default is $PWD/irs.json)
```

The port parameter is port number of web service.
```
irs web
```

Web server have some endpoints to manage irs file

Method | Endpoint | Content-Type | Info
 ------- | ------- | ------- | -------
 `POST` | `/convert` | multipart/form-data | convert irs file. will download new file.
 `GET` | `/health` | text/plain | check web server.
 `POST` | `/print` | multipart/form-data | print irs file.
 `POST` | `/validator` | multipart/form-data | validate irs file.

web page example to use irs web server:

```
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Single file upload</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="http://localhost:8208/convert" method="post" enctype="multipart/form-data">
    Format: <input type="text" name="format"><br>
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>
```

## Docker

You can run the [moov/irs Docker image](https://hub.docker.com/r/moov/irs) which defaults to starting the HTTP server.

```
docker run -p 8208:8208 moov/irs:latest
```

## Getting Started

Read through the [project docs](docs/README.md) over here to get an understanding of the purpose of this project and how to run it.

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/) | Our project documentation available online.
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
