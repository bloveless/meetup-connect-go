# Connect-go

This repository is a demonstration of building a fully type safe golang API and TypeScript react app.

It uses [connectrpc.com](https://connectrpc.com/) in order to spin up a gRPC compliant server in Go which
is also accessible to the browser without having to have gRPC behind a proxy. Technically it is running
the connect protocol which does gRPC and gRPC-web in the same server enabling connections from HTTP/1.1
and HTTP/2

## Standard development process

The easiest way to work with this repo is to change the proto definition in todo and then to run `make`.
This will do two things.

1. It will regenerate the generated files in the Go server
1. It will regenerate the generated files in the TypeScript UI

If you made any breaking changes then you'll start to see your editor light up with errors in both the
UI and the API where changes are necessary. As soon as you resolve those errors the UI and API will be
compatible again and you'll be off to the races!

## Bonus

Just for fun, and because it is dead simple in go, the web UI is embedded inside the go binary during
the build. Running `make build` will build the UI and then when the API is built it will embed those
built assets into the final go binary making both the UI and API runnable in a single binary.

Run `./connect-go` to run both the UI and API.
