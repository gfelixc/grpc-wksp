# grpc-wksp

```
git clone git@github.com:gfelixc/grpc-wksp.git
```

### Official documentation

https://developers.google.com/protocol-buffers
https://grpc.io/docs/what-is-grpc/

### Running grpc example server

```
docker run --rm -it -p 8080:8080 docker.pkg.github.com/gfelixc/grpc-wksp/flight-operator-example:latest
```


### Using a GUI Client for GRPC

- Install [BloomRPC](https://github.com/uw-labs/bloomrpc)
- Import [proto file](./server/flight_operator.proto)
- Target your server (http://localhost:8080 see [previous step](#running-grpc-example-server))
- Make your calls
