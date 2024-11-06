# IRS
**[Purpose](README.md)** | **[Configuration](CONFIGURATION.md)** | **Running**

---

## Running

### Getting Started

More tutorials to come on how to use this as other pieces required to handle authorization are in place!

- [Using docker compose](#local-development)
- [Using our Docker image](#docker-image)

No configuration is required to serve on `:8200` and metrics at `:8201/metrics` in Prometheus format.

### Docker image

You can download [our docker image `moov/irs`](https://hub.docker.com/r/moov/irs/) from Docker Hub or use this repository.

### Local Development

```
make run
```

### HTTP server

IRS runs an HTTP server at http://local.moov.io:8208 by default.

### Example using HTTP server

The HTTP server accepts JSON formatted files to convert into their PDF form. We have a few examples:

- [1099-INT](examples/1099int.json)
- [1099-MISC](examples/1099misc.json)
- [1099-OID](examples/1099oid.json)
- [1099-PATR](examples/1099patr.json)
