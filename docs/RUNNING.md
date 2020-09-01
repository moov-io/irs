# IRS
**[Purpose](README.md)** | **[Configuration](CONFIGURATION.md)** | **Running** | **[Client](../pkg/client/README.md)**

--- 

## Running

### Getting Started

More tutorials to come on how to use this as other pieces required to handle authorization are in place!

- [Using docker-compose](#local-development)
- [Using our Docker image](#docker-image)

No configuration is required to serve on `:8200` and metrics at `:8201/metrics` in Prometheus format.

### Docker image

You can download [our docker image `moov/irs`](https://hub.docker.com/r/moov/irs/) from Docker Hub or use this repository. 

### Local Development

```
make run
```

### Https service for client

https server address is https://local.moov.io:8208 in development environment.

### Example using https service 

Package github.com/moov-io/irs is a file reader and writer, should input a file as json format or ascii fire format.
Please use api endpoint of https service to creating, parsing, validating.

User can use following json files as input file for special form. 

- [1099-INT](examples/1099int.json)
- [1099-MISC](examples/1099misc.json)
- [1099-OID](examples/1099oid.json)
- [1099-PATR](examples/1099patr.json)


---
**[Next - Client](../pkg/client/README.md)**