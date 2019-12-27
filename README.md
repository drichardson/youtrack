# YouTrack for Go

[![GoDoc](https://godoc.org/github.com/drichardson/youtrack?status.svg)](https://godoc.org/github.com/drichardson/youtrack)
[![Build Status](https://travis-ci.org/drichardson/youtrack.svg?branch=master)](https://travis-ci.org/drichardson/youtrack)

Bare bones YouTrack REST API client library for Go that provides authentication
using [Permanent Token Authorization](https://www.jetbrains.com/help/youtrack/incloud/authentication-with-permanent-token.html).

The `youtrack.Api` type provides methods for making authenticated HTTP GET and POST requests to the
[YouTrack REST API](https://www.jetbrains.com/help/youtrack/incloud/youtrack-rest-api-reference.html).

There are also some convenience functions for creating issues.

Contributions welcome.

