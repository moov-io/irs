FROM debian:buster AS runtime
LABEL maintainer="Moov <support@moov.io>"
WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates \
	&& apt-get install -y pdftk \
	&& rm -rf /var/lib/apt/lists/*

COPY bin/.docker/irs /app/irs
VOLUME [ "/data", "/configs" ]

COPY ./configs/config.default.yml /configs/config.default.yml

EXPOSE 8208/tcp
EXPOSE 8209/tcp

VOLUME [ "/data", "/configs" ]

ENTRYPOINT ["/app/irs"]
CMD ["web"]
