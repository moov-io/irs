FROM debian:buster AS runtime
LABEL maintainer="Moov <support@moov.io>"
WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

COPY bin/.docker/irs /app/irs
VOLUME [ "/data", "/configs" ]

EXPOSE 8208/tcp
EXPOSE 8209/tcp

VOLUME [ "/data", "/configs" ]

ENTRYPOINT ["/app/irs"]
CMD ["web"]
