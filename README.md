# LinkedIn API Go Client

![CodeQL][12]
![Linter][13]
[![Go Reference][14]][15]
[![Go Report Card][16]][17]
[![MIT license][18]][11]

This Go package provides a thin client for making requests to LinkedIn APIs following the official [LinkedIn API documentation][01].

> :warning: This API client package is currently in beta and is subject to change. It may contain bugs, errors, or other issues that we are working to resolve. Use of this package is at your own risk. Please use caution when using it in production environments and be prepared for the possibility of unexpected behavior. We welcome any feedback or reports of issues that you may encounter while using this package.

## Versioning

`x.y.z`

- `x`:
  - `0`: under development
  - `1`: production-ready
- `y`: breaking changes
- `z`: new functionality or bug fixes in a backwards compatible manner

## Requirement

`Go 1.19+`

## Features

- [x]: Supports [Rest.li][02] protocol version 2.0.0
- [x]: Supports [LinkedIn versioned][03] APIs
- [x]: 2-legged and 3-legged OAuth2 support
- [x]: Fine-grained control over all API calls using `App` and `Session`
- [x]: Extensive documentation and [examples][04]

## License

© Mahir Hasan 2024

Released under the [MIT license][11]

[01]: https://learn.microsoft.com/en-us/linkedin/?view=li-lms-2024-04
[02]: https://linkedin.github.io/rest.li/
[03]: https://learn.microsoft.com/en-us/linkedin/marketing/versioning?view=li-lms-2024-04
[04]: _example
[11]: LICENSE
[12]: https://github.com/pilinux/linkedin/actions/workflows/codeql-analysis.yml/badge.svg
[13]: https://github.com/pilinux/linkedin/actions/workflows/golangci-lint.yml/badge.svg
[14]: https://pkg.go.dev/badge/github.com/pilinux/linkedin
[15]: https://pkg.go.dev/github.com/pilinux/linkedin
[16]: https://goreportcard.com/badge/github.com/pilinux/linkedin
[17]: https://goreportcard.com/report/github.com/pilinux/linkedin
[18]: https://img.shields.io/badge/license-MIT-brightgreen.svg
