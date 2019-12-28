# YouTrack for Go

[![GoDoc](https://godoc.org/github.com/drichardson/youtrack?status.svg)](https://godoc.org/github.com/drichardson/youtrack)
[![Build Status](https://travis-ci.org/drichardson/youtrack.svg?branch=master)](https://travis-ci.org/drichardson/youtrack)

Bare bones YouTrack REST API client library for Go that provides authentication
using [Permanent Token Authorization](https://www.jetbrains.com/help/youtrack/incloud/authentication-with-permanent-token.html).

The [youtrack.Api](https://godoc.org/github.com/drichardson/youtrack#Api)
type provides methods for making authenticated HTTP GET and POST requests to the
[YouTrack REST API](https://www.jetbrains.com/help/youtrack/incloud/youtrack-rest-api-reference.html).

There are also some convenience functions for
[creating issues](https://godoc.org/github.com/drichardson/youtrack#Api.CreateIssue).

Contributions welcome.

## Testing
The tests communicate with an actual YouTrack endpoint and therefore need credentials to run. You
can set these with the `YOUTRACK_URL` and `YOUTRACK_TOKEN` environment variables. The YouTrack
project must have at least one project with the ID (aka issue prefix or shortName) of *TP*.

To debug HTTP requests, run:

    go test -v -yt-trace

WARNING: `-yt-trace` logs the entire HTTP request, including *Authorization* token, so take care
when using it or sharing any logs generated from it.

