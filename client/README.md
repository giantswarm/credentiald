[![](https://godoc.org/github.com/giantswarm/credentiald-client?status.svg)](http://godoc.org/github.com/giantswarm/credentiald-client)
[![CircleCI](https://circleci.com/gh/giantswarm/credentiald-client/tree/master.svg?style=shield)](https://circleci.com/gh/giantswarm/credentiald-client/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/credential-client)](https://goreportcard.com/report/github.com/giantswarm/credential-client)

# credentiald-client

Go HTTP Rest client for [credentiald](https://github.com/giantswarm/credentiald) service.

## Development

### Dependencies

Dependencies are managed using [`dep`](https://github.com/golang/dep) and contained in the `vendor` directory.

### How to build

```
go build ./client
```

### How to test

```
go test ./...
```

## How to use it


```
import "github.com/giantswarm/credentiald-client/client"
```

For usage examples see: [client/client_test.go](./client/client_test.go)

## License

PROJECT is under the Apache 2.0 license. See the [LICENSE](/LICENSE) file for details.
