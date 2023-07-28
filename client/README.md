[![Go Reference](https://pkg.go.dev/badge/github.com/giantswarm/credentiald/client.svg)](https://pkg.go.dev/github.com/giantswarm/credentiald/client)

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
