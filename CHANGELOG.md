## v0.1.6 (Released 2022-02-17)

IMPROVEMENTS

- Add 1099-NEC and 2021 filing support
- chore(deps): update module github.com/moov-io/base to v0.27.5
- chore(deps): update golang.org/x/oauth2 commit hash to d3ed0bb
- chore(deps): update github.com/spf13/cobra to v1.3.0
- pdftk install script path fixed for MacOS

## v0.1.5 (Released 2021-10-15)

IMPROVEMENTS

- Fix docker crash on image startup
- chore(deps): update golang docker tag to v1.17

## v0.1.4 (Released 2021-02-08)

IMPROVEMENTS

- Add extension block for 1099ltc, 1099q, 1099r, 1099s, 1099sa, 1099sb
- Add self-test for example data, add validation logic for Combined Federal/State code
- Add document storage service with encryption

## v0.1.3 (Released 2020-08-10)

BUG FIXES

- Fix docker crash on image startup

## v0.1.2 (Released 2020-08-10)

This is the initial release of IRS which is a  CLI and HTTP server for creating, parsing, validating, and transforming IRS electronic Filing Information Returns Electronically (FIRE).

ADDITIONS

- cmd/server: Add HTTP server for parsing, printing, and validating files
- parser and validate IRS tax files
- Fuzzing support for ASCII IRS files
- OpenAPI specification for HTTP routes
